package main

import (
	"fmt"
	"io"
	"os"
)

type Console struct { // HL
	Out io.Writer
}

func (c *Console) Printf(s string, args ...interface{}) { // HL
	fmt.Fprintf(c.Out, s, args...)
}

func main() {
	c := &Console{
		Out: os.Stdout,
	}
	c.Printf("Hello, 🌍")
}
