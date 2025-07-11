package models

import (
	"strconv"
	"strings"
	"time"
)

// LineMap represents a structured log entry parsed from a raw log line.
type LineMap struct {
	IP, RemoteUser, AuthUser string
	Time                     time.Time
	ReqLine                  string
	StateCode, Size          int
	Referer, UserAgent       string
}

func (l *LineMap) GetField(fieldName *string) *string {
	switch strings.ToLower(*fieldName) {
	case "ip":
		return &l.IP
	case "remoteuser":
		return &l.RemoteUser
	case "authuser":
		return &l.AuthUser
	case "time":
		t := l.Time.Format(time.RFC1123)
		return &t
	case "reqline":
		return &l.ReqLine
	case "statecode":
		sc := strconv.Itoa(l.StateCode)
		return &sc
	case "size":
		s := strconv.Itoa(l.Size)
		return &s
	case "referer":
		return &l.Referer
	case "useragent":
		return &l.UserAgent
	default:
		return nil
	}
}
