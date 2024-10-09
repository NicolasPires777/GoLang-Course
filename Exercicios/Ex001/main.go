package main

import (
	"fmt"
	"modulo/auxiliar"
)

func main(){
	var option int
	var temperature float32

	fmt.Println("Choose your conversion:")
	fmt.Println("1 - Celsius to Fahrenheit")
	fmt.Println("2 - Fahrenheit to Celsius")
	fmt.Print("Option: ")
	_, err := fmt.Scanln(&option)

	if err != nil {
		fmt.Println("Error receiving value: ",err)
		return
	}

	if option == 1{
		fmt.Print("Enter the temperature in Celsius: ")
		_, err := fmt.Scanln(&temperature)
		if err!= nil {
			fmt.Println("Error receiving value: ",err)
			return
		}
		result := auxiliar.CelsToFir(temperature)
		fmt.Println("Temperatura converted in Fahrenheit: ",result)
		return
	} else if option == 2{
		fmt.Print("Enter the temperature in Fahrenheit: ")
		_, err := fmt.Scanln(&temperature)
		if err != nil {
			fmt.Println("Error receiving value: ",err)
			return
		}
		result := auxiliar.FirToCels(temperature)
		fmt.Println("Temperature converted in Celsius: ",result)
		return
	} else {
		fmt.Println("This option doesn't exists")
		return
	}
}