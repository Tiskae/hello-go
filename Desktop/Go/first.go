package main

import "fmt"

var num1 int = 6
const num2 int = 9

func main() {
	total := num1 + num2
	fmt.Println("The sum of the two numbers is ", total)

	test_num := 7
	test_num_pointer := &test_num

	fmt.Println((test_num_pointer))
	// Dereferencing or indirecting
	*test_num_pointer = 77
	fmt.Println(*test_num_pointer)
}	