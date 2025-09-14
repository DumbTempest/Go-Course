package main

import (
	"fmt"
)

//	func calulateTax(price float32) (float32, float32) {
//		const taxRate = 0.08
//		tax := price * taxRate
//		total := price + tax
//		return tax, total
//	}
//
//	func init() {
//		fmt.Println("init printline")
//	}
func birthday(age *int) {
	if *age < 18 {
		panic("age is less than 18")
	}
	*age = *age + 1
}

// var url = "https://gobyexample.com/"
func main() {
	// const age = 10
	// const pi float32 = 3.14
	// var m string
	// m = "hi gheys"

	// n := 42 * age // shortcut cant be used globally
	// print(m, n, pi)
	// print("\n", url)
	// _, result := calulateTax(100)
	// PrintData()
	// print("\n", data.Data_number)
	// fmt.Println(result)
	defer fmt.Println("Ended")
	defer fmt.Println("defered print")
	age := 20
	birthday(&age)
	birthday(&age)
	fmt.Println("age:", age)
}
