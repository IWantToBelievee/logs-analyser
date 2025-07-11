package models

// FilterParams holds the parameters for filtering log entries.
type FilterParams struct {
	IP         string
	RemoteUser string
	AuthUser   string
	//	Time       string
	ReqLine   string
	StateCode int
	Size      int
	Referer   string
	UserAgent string
}

func (fp *FilterParams) IsDefault() bool {
	return fp.IP == "" &&
		fp.RemoteUser == "" &&
		fp.AuthUser == "" &&
		fp.ReqLine == "" &&
		fp.StateCode == 0 &&
		fp.Size == 0 &&
		fp.Referer == "" &&
		fp.UserAgent == ""
}