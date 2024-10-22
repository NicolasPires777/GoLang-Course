package main

import "fmt"

func dia(numero int) string{
	switch numero{
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	default:
		return "Inválido"	
	}
}

func  dia2(numero int) string{
	switch{
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda"
	default:
		return "Inválido"	
	}
}

// fallthrough: Cai na proxima cláusula do switch sem analisar a condição

func main() {
	fmt.Println("Switch")

	dia:=dia(3)
	fmt.Println(dia)

	dia2 := dia2(2)
	fmt.Println(dia2)
}