package data

import "fmt"

type Duration float32

type Course struct {
	ID         int
	CourseName string
	Slug       string
	Duration   Duration
	Instructor Instructor
}

func (c Course) SignUp() bool {
	return true
}

func (c Course) String() string {
	return fmt.Sprintf("Course(ID: %d, Name: %s, Instructor: %s)", c.ID, c.CourseName, c.Instructor.Name)
}
