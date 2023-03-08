package main

import "fmt"

func main() {

	//--- les generics permettent d'avoir plusieurs types en passage de paramètre de nos  fonctions

	fmt.Println(min[float32](1, 1.5)) // paramètre de type float
	fmt.Println(min[int](1, 2))       // paramètre de type int

	//--- l'inférence nous permet de ne pas préciser le type de la variable lors de l'appel de la fonction

	pi := 3.14

	fmt.Println(max(0.7, pi))

}

func min[T int | float32](x, y T) T { // t est un type parameter prenant en compte int et float32
	return x
}

func max[F float64](x, y F) F { // le type F en inférence
	return y
}
