package Schooled

import "time"

// Student defines what all student types should implement
type Student Interface {
	Person
	Classes() []class
	Attendance() []Attenance
	Grades() []Grades
}

type Gradable interface {
	Grade()
}

type Attendable interface {
	Attendance() []Attendance
	AverageAttendance() float64
}

type Attendance struct {
	Start time.Time
	End	time.Time
	Percentage float64 // percent of class attended
}
func (a *Attendance) Duration () {
	
}

