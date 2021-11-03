package main

import (
	"time"
	"fmt"
)

type Attendee struct {
	Name string
	InterestedIn []string
	ArrivedAt time.Time
}

func main() {
	a := &Attendee{
		Name: "Gigio",
		InterestedIn: []string{
			"Coding!",
		},
		ArrivedAt: time.Now(),
	}
	fmt.Printf("%v", a)
}
