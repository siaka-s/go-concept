// lier au fichier typeStruct ou on appel le fichier de création du contact pour l'afficher
package main

type contact struct {
	name    string
	surname string
	phone   map[string]int
}

func contactcreate(name, surname string, phone1, phone2 int) contact {

	newcontact := contact{
		name:    name,
		surname: surname,
		phone: map[string]int{
			"Tel": phone1,
			"Fix": phone2,
		},
	}
	return newcontact
}
