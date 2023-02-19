package main

import "fmt"

func main() {

	w := func() {
		// affichage sans import ,envoyer à disparaitre dans le futur
		println("je suis une fonction anonyme")
	}
	// appel de la fonction anonyme
	w()

	x, y := func(x, y int) (int, int) {
		println("fonction anonyme avec retour de deux valeur dont ")
		return x, y
	}(10, 15)

	fmt.Printf("la première valeur de retour est :%v \n \bla deuxième valeur de retour est :%v", x, y)

}
