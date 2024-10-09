package main

import "fmt"

func main(){
	var number1 int8
	var number2 int8

	fmt.Print("Enter the first value: ")
	fmt.Scanln(&number1)
	fmt.Print("Enter the second value: ")
	fmt.Scanln(&number2)

	result := somar(number1,number2)
	fmt.Println("Addition result: ",result)

	sub, mult := calculos(number1,number2)
	fmt.Println("Subtraction result: ",sub)
	fmt.Println("Multiplication result: ",mult)
}

func somar(n1 int8, n2 int8) int8 {
	result := n1+n2
	return result
}

func calculos(n1, n2 int8)(int8, int8){
	r1 := n1-n2
	r2 := n1*n2
	return r1, r2
}