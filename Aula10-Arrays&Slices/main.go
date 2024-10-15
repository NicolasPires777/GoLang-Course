package main

import "fmt"

func main() {
	fmt.Println("Array e Slice")

	var array1[5] int
	array1[0] = 2
	fmt.Println(array1)

	array2 := [2]string{"pos1","pos2"}
	fmt.Println(array2)

	array3 :=[...]int{1,2,3}
	fmt.Println(array3)

	slice1 := []int{1,2,3,4,5,6,7}
	slice1 = append(slice1, 18)
	fmt.Println(slice1)

	slice2 := array3[1:3]
	fmt.Println(slice2)

	array3[2] = 9
	fmt.Println(slice2) 
}