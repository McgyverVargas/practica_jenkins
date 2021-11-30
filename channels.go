package main

import (
	"fmt"
)

func main() {
	channel := make(chan string)
	go func(channel chan string){
		for{
			var name string
			fmt.Scanln(&name)
			channel <- name
		}
	}(channel)
	/* msg := <- channel
	fmt.Println("Aqui recibo " + msg)
	msg = <- channel
	fmt.Println("Aqui recibo 2 " + msg) */

	for i := 0; i < 3; i++ {
		msg := <- channel
	    fmt.Println(i," Aqui recibo " + msg)
	}
}