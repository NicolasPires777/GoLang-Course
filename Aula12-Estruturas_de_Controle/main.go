package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")

	numero := 10

	if numero > 15{
		fmt.Println("Maior que quinze")
	} else {
		fmt.Println("Menor ou igual que quinze")
	}

	if outro := numero; outro>0{
		fmt.Println("Outro é maior que zero")
	}

	//else if (caso precise)
}