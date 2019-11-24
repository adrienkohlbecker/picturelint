package validators

//go:generate stringer -type=Status

type Status int

const (
	StatusUndefined Status = iota
	StatusSuccess
	StatusFailed
	StatusSkipped
)
