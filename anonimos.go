package main

import "fmt"

func main() {
	pru := Tutor{Human{name:"Estarosa"}}
	fmt.Println(pru.name)
	fmt.Println(pru.Human.name)
}

type Human struct {
	name string
}

type Tutor struct {
	Human
}