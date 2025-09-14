package main

import (
	"fmt"
	"go/FEM/data"
)

func main() {
	max := data.Instructor{}
	max.Name = "Rajat"
	// ram := data.NewInstructor("Ram", "meow")

	goCourse := data.Course{ID: 2, CourseName: "Go fundamentals", Instructor: max}
	// println(max.Print())
	// println(ram.Print())
	fmt.Printf("%+v", goCourse)

	swiftWS := data.NewWorkshop("Swift Fundamentals", max)
	fmt.Printf("%+v", swiftWS)

	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = swiftWS

	for _, course := range courses {
		fmt.Println(course.SignUp())
	}

}
