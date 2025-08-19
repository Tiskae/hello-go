package main

import (
	"fmt"
	"strings"
)

func main() {
	var best_five [5]string
	best_five[0] = "Lamborghini Aventador SVJ"
	best_five[1] = "Pagani Zonda"
	best_five[2] = "BMW M4 F82"
	best_five[3] = "Dodge Challenger Demon"
	best_five[4] = "Acura NSX"

	// var slice = []int{0, 1, 2, 3, 4, 5}
	var best_slice []string = best_five[0:3]
	best_slice[2] = "Mercedes CLA 🤢"
	new_best_slice := append(best_slice, "Bentley Continental GT")

	fmt.Println(best_five)
	fmt.Println(best_slice)
	fmt.Println(new_best_slice)

	// Array & Slice literal
	var arr_literal = [3]bool{true, false, true}
	var slice_literal = []bool{true, false, false, true}

	fmt.Println(arr_literal, slice_literal)

	// Slice defaults - all are equivalent
	slice := best_five[0:5]
	slice = best_five[0:]
	slice = best_five[:5]
	slice = best_five[:]

	fmt.Println(slice)

	// Slice length & capacity
	length := len(best_slice)
	capacity := cap(best_slice)

	fmt.Println(length, capacity)

	// Slice - creation with inbuilt "make" func
	var make_slice = make([]int, 0, 10)
	fmt.Println(make_slice)

	// Slices of slices
	var board = [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[1][0] = "O"
	board[1][1] = "X"
	board[2][0] = "O"
	board[2][2] = "X" // X wins

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
