package data

import "time"

type Workshop struct {
	Course
	Instructor
	Date time.Time
}

func (c Workshop) SignUp() bool {
	return true
}
func NewWorkshop(name string, instructor Instructor) Workshop {
	w := Workshop{}
	w.Name = name
	w.Course.Instructor = instructor
	w.Date = time.Now()
	return w
}
