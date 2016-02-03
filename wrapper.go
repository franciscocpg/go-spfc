// Package service provides Start, Status and Stop functions
package service

type StatusResponse struct {
	Running bool
	PID     int
}

// Status show the status for a given service name (s)
func Status(s string) (StatusResponse, error) {
	return status(s)
}
