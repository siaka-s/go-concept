package main

import "fmt"

func array() {

	// déclaration
	// var tab [5]int

	// boucle de remplissage
	// for i := 0; i < len(tab); i++ {

	// 	tab[i] = i
	// }
	// fmt.Println(tab)

	//Déclaration et assignation sans nombre de valeur
	// tab := [5]int{1, 2, 3, 4, 5}

	//Déclaration et assignation sans nombre de valeur
	tab := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, v := range tab {

		fmt.Println(index, ":", v)

	}
}

func main() {
	array()

}
