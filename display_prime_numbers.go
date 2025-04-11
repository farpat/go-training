package main

import (
	"errors"
)

func getPrimeNumbers(n int) []int {
	var primeNumbers []int

	for index := 0; index < n; index++ {
		isPrimeNum, err := isPrime(index)
		if err == nil && isPrimeNum {
			primeNumbers = append(primeNumbers, index)
		}
	}

	return primeNumbers
}

func isPrime(n int) (bool, error) {
	if n < 1 {
		return false, errors.New("n must be greater than 0")
	}

	if n <= 2 {
		return true, nil
	}

	if n%2 == 0 {
		return false, nil
	}

	for index := 3; index < (n / 2); index++ {
		if n%index == 0 {
			return false, nil
		}
	}

	return true, nil
}
