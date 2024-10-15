package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main(){
	fmt.Println("Escrevi no main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("devbookgmail.com")
	fmt.Println(erro)
}

//"go mod tidy" para limpar dependencia não utilizadas 