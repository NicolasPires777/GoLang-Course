package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string] string { //Chaves devem ser todas do mesmo tipo e valores todos do mesmo tipo
		"nome":"Pedro",
		"Sobrenome":"Silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome":{
			"primeiro":"João",
			"ultimo":"Pedro",
		},
	}

	fmt.Println(usuario2["nome"]["primeiro"])
	delete(usuario2,"nome")
	fmt.Println(usuario2)
}