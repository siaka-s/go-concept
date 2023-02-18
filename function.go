package main

import (
	"fmt"
)

func SayName(name string) string {

	message := fmt.Sprintf("Mon nom est %v", name)
	// valeur de retour
	return message

}

func main() {

	fmt.Println(SayName("siaka-s"))

}
