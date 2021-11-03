package main

import "fmt"

// FI OMIT
type Finder interface {
	// Position returns latitude, longitude and the precision of the
	// reading in meters.
	Position() (float64, float64, int) // HL
}
// FO OMIT

// II OMIT
type Item struct {
	Finder // HL
	Id string
}

func NewItem(id string, finder Finder) Item {
	return Item{Finder: finder, Id: id}
}
// IO OMIT

func (i Item) printPosition() {
	latitude, longitude, precision := i.Position()
	fmt.Printf("%q: %f, %f (± %dm)\n", i.Id, latitude, longitude, precision)
}

// FMNI OMIT
type FindMyNetwork struct {}

func (n FindMyNetwork) Position() (float64, float64, int) { // HL
	return 46.479186, 11.332571, 150
}

type GPS struct {}

func (s GPS) Position() (float64, float64, int) { // HL
	return 46.479186, 11.332571, 1
}

// FMNO OMIT

// MI OMIT
func main() {
	NewItem("abc", GPS{}).printPosition()
}
// MO OMIT
