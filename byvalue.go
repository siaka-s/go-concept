package main

import "fmt"

func maj(num int) int {
	num = 5
	// retour d'une valeur pour la fonction pour permettre l'assignation dans main
	return num
}

// mise a jour d'un élement d'une maps
func majmap(item map[string]int) {

	item["bonbon"] = 200

}

func main() {

	num := 10

	maj(num)

	fmt.Println(num) // 10

	// on réassigne la valeur de la variable pour permettre mla modification num (valeur)
	num = maj(num)

	fmt.Println(num) // 5

	listecourse := map[string]int{

		"sel":   500,
		"huile": 1500,
		"pain":  150,
	}
	majmap(listecourse) // bonbon est ajouté

	fmt.Println(listecourse)

}
