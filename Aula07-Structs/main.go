package main

import "fmt"

type usuario struct{
	nome string
	idade uint8
	local endereco
}

type endereco struct{
	Street string
	Number uint
}

func main(){
	local := endereco{"Rua dos lobo", 456}
	var person1 usuario
	person1.nome = "Davi"
	person1.idade = 25
	person1.local = local
	fmt.Println(person1)

	person2 := usuario{"Nicolas",20, local}
	fmt.Println(person2)
}