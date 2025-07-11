package src

import (
	"context"
	"fmt"
	"logs-analyser/pkg/models"
	"strings"
	"sync"
)

// FilterFunc is a function type that takes a LineMap and returns a boolean
type FilterFunc func(models.LineMap) bool

// Filter processes a channel of LineMap entries, applying a series of filters
func Filter(input *chan models.LineMap, output *chan models.LineMap, filters []FilterFunc, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	// Return without closing the output channel if no filters are provided
	if len(filters) == 0 {
		return
	}
	defer close(*output)

	select {
	case <-ctx.Done():
		fmt.Println("Scraper shutting down")
		return
	default:
	OUTER:
		for l := range *input {
			for _, f := range filters {
				if !f(l) {
					continue OUTER
				}
				*output <- l
			}
		}
	}
}

// SelectFilters constructs a slice of FilterFunc based on the provided FilterParams
func SelectFilters(params models.FilterParams) []FilterFunc {
	filters := []FilterFunc{}

	if params.IP != "" {
		filters = append(filters, filterByIP(params.IP))
	}
	if params.RemoteUser != "" {
		filters = append(filters, filterByRemoteUser(params.RemoteUser))
	}
	if params.AuthUser != "" {
		filters = append(filters, filterByAuthUser(params.AuthUser))
	}
	if params.ReqLine != "" {
		filters = append(filters, filterByReqLine(params.ReqLine))
	}
	if params.StateCode != 0 {
		filters = append(filters, filterByStateCode(params.StateCode))
	}
	if params.Size != 0 {
		filters = append(filters, filterBySize(params.Size))
	}
	if params.Referer != "" {
		filters = append(filters, filterByReferer(params.Referer))
	}
	if params.UserAgent != "" {
		filters = append(filters, filterByUserAgent(params.UserAgent))
	}

	return filters
}

func filterByIP(ip string) FilterFunc {
	return func(l models.LineMap) bool {
		return l.IP == ip
	}
}

func filterByRemoteUser(remoteUser string) FilterFunc {
	return func(l models.LineMap) bool {
		return l.RemoteUser == remoteUser
	}
}

func filterByAuthUser(authUser string) FilterFunc {
	return func(l models.LineMap) bool {
		return l.AuthUser == authUser
	}
}

func filterByReqLine(reqLine string) FilterFunc {
	return func(l models.LineMap) bool {
		return strings.Contains(l.ReqLine, reqLine)
	}
}

func filterByStateCode(code int) FilterFunc {
	return func(l models.LineMap) bool {
		return l.StateCode == code
	}
}

func filterBySize(size int) FilterFunc {
	return func(l models.LineMap) bool {
		return l.Size == size
	}
}

func filterByUserAgent(userAgent string) FilterFunc {
	return func(l models.LineMap) bool {
		return strings.Contains(l.UserAgent, userAgent)
	}
}

func filterByReferer(referer string) FilterFunc {
	return func(l models.LineMap) bool {
		return strings.Contains(l.Referer, referer)
	}
}

// func filterByTimeRange(start, end time.Time) FilterFunc {
// 	return func(l models.LineMap) bool {
// 		return l.Time.After(start) && l.Time.Before(end)
// 	}
// }
