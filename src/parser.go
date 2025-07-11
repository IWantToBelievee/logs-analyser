package src

import (
	"context"
	"fmt"
	"logs-analyser/pkg/models"
	"regexp"
	"strconv"
	"sync"
	"time"
)

// Parser reads raw log lines from rawLogLines channel,
// parses them according to the Apache Combined Log Format,
// and sends structured LineMap objects to parsedLines channel.
//
// Parameters:
//
//	rawLogLines: A pointer to a channel that provides raw log lines as strings.
//	parsedLogLines: A pointer to a channel where parsed 'LineMap' objects will be sent.
func Parser(rawLogLines *chan string, parsedLogLines *chan models.LineMap, wg *sync.WaitGroup, ctx context.Context) {
	parseIP := regexp.MustCompile(`(\d+)\.(\d+)\.(\d+)\.(\d+)`)
	parseRemoteUser := regexp.MustCompile(`^\d{1,3}(?:\.\d{1,3}){3}\s+([^\s]+)`)
	parseAuthUser := regexp.MustCompile(`^\d{1,3}(?:\.\d{1,3}){3}\s+[^\s]+\s+([^\s]+)`)
	parseTime := regexp.MustCompile(`\[(\d{2})/([a-zA-Z]+)/(\d{4}):(\d{2}):(\d{2}):(\d{2}) ([+-]\d{4})\]`)
	parseReqLine := regexp.MustCompile(`"(GET|POST|PUT|DELETE|HEAD|OPTIONS|PATCH|TRACE|CONNECT)\s+([^\s]+)\s+(HTTP|HTTPS)/(\d+)\.(\d+)"`)
	parseStateCode := regexp.MustCompile(`"\s+(\d{3})\s+\d+`)
	parseSize := regexp.MustCompile(`\s+(\d+)\s+"`)
	parseReferer := regexp.MustCompile(`"[^"]+"\s+\d+\s+\d+\s+"([^"]+)"`)
	parseUserAgent := regexp.MustCompile(`"([^"]+)"\s*$`)

	const timeLayout = "02/Jan/2006:15:04:05 -0700"

	defer wg.Done()
	defer close(*parsedLogLines)

	select {
	case <-ctx.Done():
		fmt.Println("Parser shutting down")
		return
	default:
		// Iterate over each raw log line received from the input channel.
		for l := range *rawLogLines {
			timeStr := parseTime.FindString(l)
			// Remove the surrounding square brackets.
			timeStr = timeStr[1 : len(timeStr)-1]

			// Parse the timestamp string into a time.Time object.
			t, err := time.Parse(timeLayout, timeStr)
			if err != nil {
				panic(err)
			}

			// Populate the LineMap struct with parsed data.
			lineMap := models.LineMap{
				IP:         parseIP.FindString(l),
				RemoteUser: parseRemoteUser.FindStringSubmatch(l)[1],
				AuthUser:   parseAuthUser.FindStringSubmatch(l)[1],
				Time:       t,
				ReqLine:    parseReqLine.FindString(l),
				StateCode: func() int {
					stateCode, err := strconv.Atoi(parseStateCode.FindStringSubmatch(l)[1])
					if err != nil {
						panic(err)
					}
					return stateCode
				}(),
				Size: func() int {
					size, err := strconv.Atoi(parseSize.FindStringSubmatch(l)[1])
					if err != nil {
						panic(err)
					}
					return size
				}(),
				Referer:   parseReferer.FindStringSubmatch(l)[1],
				UserAgent: parseUserAgent.FindString(l),
			}
			// Send the fully parsed LineMap object to the output channel.
			*parsedLogLines <- lineMap
		}
	}
}
