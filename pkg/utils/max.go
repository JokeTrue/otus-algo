package utils

func FindMax(array []int) int {
	max := array[0]

	for _, value := range array[1:] {
		if value > max {
			max = value
		}
	}

	return max
}
