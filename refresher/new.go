package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World!")

	// This is a single line comment
	/*
	 * This is a multi-line comment
	 */

	// var variablename type = value
	var strVar string = "String variable"

	var (
		userAge     int    = 24
		userBalance int    = 300
		userName    string = "Abdullahi Al Qadrashi"
	)

	const newConstWithNoType = "Yayy!"

	fmt.Println(userAge, userBalance, userName, "\n", newConstWithNoType)
	fmt.Println(strVar)
	fmt.Println(10, 20)

	// %s string, %d decimal, %t boolean, %f float
	fmt.Printf("My name is %s and the type is %T. My account balance is %+d \n", "Tiskae", "Tiskae", 23_000)
	fmt.Println(math.Abs(-10000))

	var arr1 = [5]int{1, 2, 3}
	var arr2 = []int{1, 2, 3}
	cars := [7]string{"BMW", "Merecdes", "Porsche", "Audi", "Volkswagen", "Koenigesgg", "Lamborghini"}

	fmt.Println(arr1, arr2)
	fmt.Println(cars)
}
