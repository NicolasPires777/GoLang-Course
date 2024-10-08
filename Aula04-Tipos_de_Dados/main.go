package main

import (
	"errors"
	"fmt"
)

func main(){
	var num1 int16 = 100
	fmt.Println(num1)

	var flo1 float32 = 1.34
	fmt.Println(flo1)

	char := 'M'
	fmt.Println(char)

	var booleano1 bool = true
	booleano2 := false
	fmt.Println(booleano1,booleano2)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}

// Tipos de inteiro: int(usa arquitetura do PC como base), int8, int16, int32, int64, uint(inteiro sem sinal)
// Tipos de Float: float32, float64


//INT32 = Rune
//UINT8 = byte