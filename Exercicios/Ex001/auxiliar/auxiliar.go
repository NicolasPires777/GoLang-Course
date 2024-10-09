package auxiliar

func CelsToFir(n1 float32) float32 {
	var Fahrenheit float32 = n1*9/5+32
	return Fahrenheit
}

func FirToCels(n1 float32) float32{
	var Celsius float32 = (n1-32)*5/9
	return Celsius
}