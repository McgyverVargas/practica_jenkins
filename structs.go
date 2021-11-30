package main

import "fmt"

func main() {
	pru := User{edad : 15, nombre : "Pedro"}
	pru.nombre = "Lucas"
	pru.apellido = "Perez"
	fmt.Println(pru.nombre)
	pru.set_Nombre("Juan")
	fmt.Println(pru.name_completo())
}

type User struct {
	edad int
	nombre, apellido string
}

func (usuario User) name_completo() string {
	return usuario.nombre + " " + usuario.apellido
}

func (usuario *User) set_Nombre(n string) {
	usuario.nombre = n
}