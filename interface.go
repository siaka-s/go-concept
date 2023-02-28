package main

import "fmt"

type Car interface {
	vitesseMax() int

	kilometrage() int
}

type Kia struct {
	model string

	color string

	prix int
}

type Ford struct {
	model string

	color string

	prix int
}

func main() {

	kia := Kia{
		model: "sportage phase 2",
		color: "white",
		prix:  21000,
	}
	printCarInfo(&kia)

	ford := Ford{
		model: "focus phase 1",
		color: "red",
		prix:  15000,
	}
	printCarInfo(&ford)

}

func printCarInfo(c Car) {
	fmt.Println("la voiture peut atteindre", c.kilometrage(), "km avec une vitesse max estimé à", c.vitesseMax(), "km/h")
}

func (k *Kia) vitesseMax() int {
	return 260
}
func (k *Kia) kilometrage() int {
	return 1000
}
func (f *Ford) vitesseMax() int {
	return 230
}
func (f *Ford) kilometrage() int {
	return 750
}
