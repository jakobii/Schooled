package school

import (
	"github.com/google/uuid"
	"time"
)

// Session
type Session struct {
	ID         uuid.UUID
	Start, End time.Time
	Teachers   []Teacher
	Attendance []Attender
}

// Duration gets the duration of a class session.
func (s *Session) Duration() (d time.Duration) {
	return d
}
