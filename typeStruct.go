package main

import "fmt"

type personnel struct {
	id int

	name string

	surname string

	adult bool
}

func main() {

	personnel1 := personnel{1, "krépin", "charles", false} // obligation de spécifier tout les type

	personnel2 := personnel{ // assignation d'une valeur par default si le type est pas spécifié

		// id:   2,       // 0

		name: "kouadio", // ""

		surname: "kevin", // ""

		adult: true, // false

	}

	fmt.Println(personnel1)
	fmt.Println(personnel2)

}
