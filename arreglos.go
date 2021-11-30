package main

import "fmt"

func main() {
	arreglo := [4]string {"Mama","Papa","Hijo"}
	arreglo[3] = "Nieto"
	for i := 0; i < len(arreglo); i++ {
		fmt.Println(arreglo[i])		
	}
	fmt.Println(arreglo)
	matriz := []int{1,2,3,4}
	fmt.Println(matriz)
	slice := arreglo[1:2]
	fmt.Println(slice)
	makes := make([]int,4,5)
	fmt.Println(makes)
	fmt.Println(len(makes))
	fmt.Println(cap(makes))
	makes = append(makes,3)
	fmt.Println(makes)
	copia := make([]int,4)
	copy(copia,matriz)
	fmt.Println(copia)
}