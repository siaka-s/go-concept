package main

import (
	"fmt"
)

func main() {
	// Declaration de variable
	var (
		name string
		age  int
	)
	// initialisation
	age = 15

	name = "julien"

	// condition
	if age < 18 {
		// formatage Println
		fmt.Println(name, "est mineur car il y a moins de", age, "ans")

	} else if age >= 18 {
		//  formatage avec Printf
		fmt.Printf("%v est majeur car il a plus de %v ans", name, age)

	}

}
