package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main() {
	fmt.Println("Herança")
	p1 := pessoa{"João", "Pedro", 20, 178}
	fmt.Println(p1)
	a1 := estudante{p1,"ADS","FMP"}
	fmt.Println(a1)

	fmt.Println("O estudante se chama",a1.nome)
}