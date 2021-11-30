package main

import (
	/* "io/ioutil" */
	"fmt"
	"bufio"
	"os"
)

func main() {
	/* file_data,err := ioutil.ReadFile("./helps.txt")
	if(err != nil){
		fmt.Println("Hubo error")
	}
	fmt.Println(string(file_data)) */
	valida := readFile()
	fmt.Println(valida)	
}

func readFile() bool{
	abrir_data,err2 := os.Open("./helps2.txt")
	defer func(){
		abrir_data.Close()
		fmt.Println(recover())
	}()	
	if(err2 != nil){
		fmt.Println("Hubo error")
		panic(err2)
	}
	scanner := bufio.NewScanner(abrir_data)
	for(scanner.Scan()){
		linea := scanner.Text()
		fmt.Println("aqui es " + linea)
	}
	return true
}