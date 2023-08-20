package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	var word_map = make(map[string]int)
	for _, v := range strings.Fields(s) {
		if _, ok := word_map[v]; ok {
			word_map[v]++
		} else {
			word_map[v] = 1
		}
	}
	return word_map
}

func main() {
	var result = WordCount("I am learning Go , let's Go !")
	fmt.Println(result)
}
