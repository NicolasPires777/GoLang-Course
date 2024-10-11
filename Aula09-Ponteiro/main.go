package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var var1 int8 = 10
	var var2 int8 = var1

	fmt.Println(var1,var2)

	var var3 int8 = 100
	var ponteiro *int8

	ponteiro = &var3

	fmt.Println(var3, ponteiro)

	var3 = 15

	fmt.Println(var3, *ponteiro)
}