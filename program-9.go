package main

import "fmt"

func main() {
	fmt.Print("Convertir decimal a binario: \n")
	fmt.Print("Ingrese un número decimal: ")
	var decimalNumber int64
	fmt.Scanln(&decimalNumber)

	fmt.Printf("%b", decimalNumber)
}
