package main

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
)

func main() {
	/* var nombrevariable string
	var x,y,z int
	var valida bool */
	x := strconv.Itoa(23)
	y, _ := strconv.Atoi(x)
	bandera := false
	//nombre := "Prueba"
	fmt.Println("Este numero es " + x)
	fmt.Println(20 + y)
	fmt.Println("Esto es ", bandera)

	/* var name string
	fmt.Println("Digite un nombre: ")
	fmt.Scanf("%s\n",&name)
	fmt.Printf("El nombre digitado es %s\n",name) */

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese un valor: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El valor es ",text)
	}
}