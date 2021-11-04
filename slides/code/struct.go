package main

import "fmt"

// I OMIT
type GPS struct {}

func (s GPS) Position() (float64, float64, int) {
	return 46.479186, 11.332571, 1
}

type Item struct {
	GPS
	Id string
}

func NewItem(id string) Item {
	return Item{
		GPS: GPS{},
		Id: id,
	}
}

func main() {
	i := NewItem("abc")
	latitude, longitude, precision := i.GPS.Position()
	fmt.Printf("%f, %f (± %dm)\n", latitude, longitude, precision)
}
// O OMIT
