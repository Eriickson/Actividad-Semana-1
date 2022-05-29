package main

import "fmt"

func main() {
	fmt.Print("Restar 2 números: \n")
	fmt.Print("Ingrese el primer número: ")
	var firstNumber int8
	fmt.Scanln(&firstNumber)

	fmt.Print("Ingrese el segundo número: ")
	var secondNumber int8
	fmt.Scanln(&secondNumber)

	fmt.Print("El resultado es: ", firstNumber-secondNumber)
}
