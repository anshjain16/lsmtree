package util

type Status string

const (
	STATUS_OK Status = "OK"
	STATUS_NOT_OK Status = "NOT_OK"
	STATUS_UNKNOWN Status = "UNKNOWN"
)

func(s Status) String() string {
	return string(s)
}
