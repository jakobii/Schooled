package school

import (
	"github.com/google/uuid"
	"time"
)


// Attender interface is used to calculate standard attendance information.
// Attended should simply state whether or not a student attended class
// duration should return the amount of time spent in class.
type Attender interface {
	Attended() bool
  Duration() time.Duration
}

// PercentAttendance represented a students attendance to a class session.
// Storing attendance information as a percentage will make dealing
// with clarical scheduling errors more dynamic. If a teacher says
// that a student showed up for 20% of the class day and the
// schedule was incorrectly stated to be 2 minutes instead of 2
// hours, the schedule can be updated and attendance records will
// still be accurate as long as the schedule accurately reflects
// historical events.
type PercentAttendance struct {
	Session *Session
  Student *Student
	ID      uuid.UUID
	Percent float64 // percent of class attended
}

// Attended return true if a student attended a class session
// for any amount of time and false if they never showed up.
func (a *PercentAttendance) Attended() (b bool) {
	if a.Percent > 0 {
		return true
	}
	return false
}

// TimeAttendance is used to capture the duration of time a student spent in a
// class session. It is less flexalble the PercentAttendance but can has the
// advantage of being explicitly more accurate.
type TimeAttendance struct {
	Session
	ID   uuid.UUID
	Time time.Duration // percent of class attended
}

// Attended return true if a student attended a class session
// for any amount of time and false if they never showed up.
func (a *TimeAttendance) Attended() (b bool) {
	if a.Time > 0 {
		return true
	}
	return false
}

type BoolAttendance struct {
	Session
	ID   uuid.UUID
	Bool bool
}

// Attended return true if a student attended a class session
// for any amount of time and false if they never showed up.
func (a *BoolAttendance) Attended() (b bool) {
	return a.Bool
}
