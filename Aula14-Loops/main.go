package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")
	
	i:=0

	for i < 1 {
		i++
		fmt.Println(i)
		time.Sleep(time.Second)
	}

	for j:= 0; j<1; j++ { // j += 2 
		fmt.Println(j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"Joao","Davi","Lucas"}

	for _, nome := range nomes{
		fmt.Println(nome)
	}

	for indice, letra := range "Palavra"{
		fmt.Println(indice, string(letra))
	}

	usuario := map[string] string{
		"nome": "leonardo",
		"Sobre": "Silva",
	}

	for a, i := range usuario{
		fmt.Println(a, i)
	}

	
}