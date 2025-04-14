// Package divisors implements the divisors exercise
package divisors

// GetDivisors returns all divisors of a number
func GetDivisors(n int) []int {
	var divisors []int
	for i := 1; i <= (n / 2); i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}

	divisors = append(divisors, n)
	return divisors
}
