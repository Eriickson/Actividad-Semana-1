// Simple Program to Check Entered Number is Even or Odd

package main

import "fmt"

func main() {
	fmt.Print("Verificar si un número es par o impart : \n")
	fmt.Print("Ingresa un número : ")
	var n int
	fmt.Scanln(&n)

	if n%2 == 0 {
		fmt.Println(n, "es un número par")
	} else {
		fmt.Println(n, "es un número impar")
	}
}
