package primes

import "fmt"

func getPrimeNumbers(n int) []int {
	var primeNumbers []int

	for index := 0; index < n; index++ {
		if isPrime(index) {
			primeNumbers = append(primeNumbers, index)
		}
	}

	return primeNumbers
}

func isPrime(n int) bool {
	if n%2 == 0 {
		return false
	}

	for index := 3; index < (n / 2); index++ {
		if n%index == 0 {
			return false
		}
	}

	return true
}

func main() {
	max := 50

	fmt.Printf("Prime numbers up to %d:\n", max)
	fmt.Println(getPrimeNumbers(max))
}
