package main

import "fmt"

func getDivisorsCount(n int) int {
	divisorsCount := 0

	for i := 1; i <= (n / 2); i++ {
		fmt.Println("Checking", i)
		if n%i == 0 {
			divisorsCount++
		}
	}

	return divisorsCount + 1
}
