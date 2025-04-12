package main

func getPrimeNumbers(n uint) []int {
	var primeNumbers []int

	for index := uint(0); index < n; index++ {
		if isPrime(index) {
			primeNumbers = append(primeNumbers, int(index))
		}
	}

	return primeNumbers
}

func isPrime(n uint) bool {
	if n <= 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	for index := uint(3); index < (n / 2); index++ {
		if n%index == 0 {
			return false
		}
	}

	return true
}
