package main

import "fmt"

func main() {
	for i := 1; i < 10; i++{
		fmt.Println("Hola a todos ", i)
	}
	j := 1
	for j < 10 {
		fmt.Println("Hola a todos otra vez ", j)
		j++
	}
}