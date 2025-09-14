package data

import "fmt"

var Countries [10]string

func init() {
	Countries[0] = "Argentina"
	Countries[1] = "Brazil"
	Countries[8] = "USA"
	fmt.Println("init from data package")
}
