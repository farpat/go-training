package main

func countBits(n uint) int {
	count := 0

	for n > 0 {

		if n%2 != 0 {
			count++
		}

		n /= 2
	}

	return count
}
