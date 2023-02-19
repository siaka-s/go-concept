package main

import "fmt"

func point(y *int) { // y pointe vers l'adresse memoire de x

	*y += 10
}

func main() {

	x := 10

	point(&x) // récupération de l'adresse mémoire

	fmt.Println(x)

}
