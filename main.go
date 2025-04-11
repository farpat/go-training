package main

import "fmt"

func main() {
	max := 50

	fmt.Printf("Prime numbers up to %d:\n", max)
	fmt.Println(getPrimeNumbers(max))

	fmt.Println(displaySquare(0))
	fmt.Println("")
	fmt.Println(displaySquare(1))
	fmt.Println("")
	fmt.Println(displaySquare(2))
	fmt.Println("")
	fmt.Println(displaySquare(3))
	fmt.Println("")
	fmt.Println(displaySquare(4))
}
