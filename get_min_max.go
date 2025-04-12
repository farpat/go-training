package main

func getMinMax(numbers []int) [2]int {
	min := numbers[0]
	max := numbers[0]

	for _, number := range numbers {
		if number < min {
			min = number
		}

		if number > max {
			max = number
		}
	}

	return [2]int{min, max}
}
