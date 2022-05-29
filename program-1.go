package main

import "fmt"

// Buscar el factorial de un numero
func Factorial(n int) int {

	if n <= 1 {
		return n
	}

	return n * Factorial(n-1)

}

func main() {

	fmt.Println(Factorial(1))
	fmt.Println(Factorial(2))
	fmt.Println(Factorial(3))
	fmt.Println(Factorial(4))
	fmt.Println(Factorial(5))

}
