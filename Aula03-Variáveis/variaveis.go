package main

import "fmt"

func main(){
	var var1 string = "Isto é uma string"
	var2 := "Variavel 2"
	fmt.Println(var1)
	fmt.Println(var2)

	var (
		var3 string = "3"
		var4 string = "2"
	)
	fmt.Println(var3,var4)

	var5,var6 := "5", "6"
	fmt.Println(var5,var6)

	const cons1 string = "valor 1"
	fmt.Println(cons1)

	var5, var6 = var6, var5 // Inverte o valor das variáveis
}