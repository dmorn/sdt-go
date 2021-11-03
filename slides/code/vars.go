package main

import (
	"fmt"
	"time"
)

func printValType(v interface{}) {
	fmt.Printf("%T: %v\n", v, v)
}

// I OMIT
func main() {
	// Long version
	var lang string
	lang = "Go"

	version := "go1.17.2 darwin/amd64"
	now := time.Now()
	i := 7 + 1
	f := 9.0 / 2

	printValType(lang)
	printValType(version)
	printValType(now)
	printValType(i)
	printValType(f)

	printValType("anything!")
}
// O OMIT
