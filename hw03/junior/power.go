package main

func Power(number int, power int) int {
	result := 1

	for i := 0; i < power; i++ {
		result *= number
	}

	return result
}
