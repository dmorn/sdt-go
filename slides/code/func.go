package main

import "fmt"

// invisible outside main package
func addMore(nums ...int) int { // HL
	var result int
	for _, v := range nums {
		result += v
	}
	return result
}

// visibile outside main
func Add(a, b int) int { // HL
	return addMore(a, b)
}

func main() {
	sum := Add(1, 99)
	fmt.Println(sum)
}
