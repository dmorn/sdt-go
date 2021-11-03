package main

import "fmt"

func main() {
	i := 42

	p := &i // HL
	*p = 21 // HL
	fmt.Println(i)
}
