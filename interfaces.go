package main

import "fmt"

func main() {
	adminis := []User{Admin{Name:"Estarosa"},Admin{Name:"Meliodas"}}
	for _, admin := range adminis{
		fmt.Println(auth(admin))
	}
}

type User interface {
	Permisos() int
	Nombres() string
}

type Admin struct {
	Name string
}

func (this Admin) Permisos() int{
	return 3
}

 func (this Admin) Nombres() string{
	 return this.Name
 }

 func auth(this User) string{
	 permises := this.Permisos()
	 name := this.Nombres()
	 if(permises > 2){
		 return name + " Tiene permisos de Admin"
	 }
	 return name + " No Tiene permisos de Admin"
 }