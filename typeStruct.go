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

	// Appel de la fonction contactcreat contenu dans contact.go

	newc := contactcreate("siaka", "siahoué", 07222222, 0722266666) // lier au fichier contact qui cré un nouveau contact

	fmt.Println(newc)

	fmt.Printf("le nom est : %v ", newc.name) // récupéré une valeur de mon type contact

}

// exécution du code -> go run typeStruct.go contact.go
