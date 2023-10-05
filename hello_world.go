package main

import "fmt"

func main() {

	//Hola mundo
	/* Hola mundo
	en varias lineas*/
	fmt.Println("Hola, Go")
	//variables
	// var "name" "type"
	var myString string = "Esto es una cadena de texto"
	fmt.Println(myString)
	myString = "aqui cambia el valor de la cadena de texto"
	fmt.Println(myString)
	var myString2 = "Hola esto es string2 sin tipado"
	fmt.Println(myString2)

	//si se declara como string se muere como string
	var myInt int = 7
	fmt.Println(myInt)
	//int8 int16 *int 32 int64

	myInt = myInt + 4
	fmt.Println(myInt)
	fmt.Println(myInt - 1)
	fmt.Println(myInt + 4)

}
