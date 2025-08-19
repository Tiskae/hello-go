package main

import (
	"fmt"
	"math"
	"math/rand"
)

func multipleArgsFunc(args... int)  {
	fmt.Println(args)
}

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

	var arr1 = []int{1, 2, 3}
	var arr2 = []int{1, 2, 3}
	cars := [7]string{"BMW", "Merecdes", "Porsche", "Audi", "Volkswagen", "Koenigesgg", "Lamborghini"}
	emptyArr := [5]string{1: "Nothing", 4: "Void"}
	newArr := append(arr1, arr2...)

	fmt.Println(arr1, arr2)

	cars[4] = "Bentley"
	fmt.Println(cars, emptyArr)
	fmt.Println(newArr, len(newArr), cap(newArr))
	fmt.Println(5 | 2)
	fmt.Println(5 / (math.Pow(7, 0) - math.Pow(1002, 0)))
	fmt.Println(5 << 2)

	// Grade class calculator
	var score int = rand.Intn(101);
	var grade string

	if score >= 70 {
		grade = "A"
	} else if score >= 60 {
		grade = "B"
	} else if score >= 50 {
		grade = "C"
	} else if score >= 45 {
		grade = "D"
	} else if score >= 40 {
		grade = "E"
	} else if score < 40 {
		grade = "F"
	} else {
		grade = "Invalid score"
	}

	fmt.Printf("Your grade is %s \n", grade)

	nums := []int{1,2,3,4,5,6,7,8,9, 10}

	for i, j := range nums {
		fmt.Println(i, j)
	}



	multipleArgsFunc(8, 2, 3, 4)
}
