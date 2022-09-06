package sort

import "otus-algo/pkg/utils"

func SelectionSort(array []int) []int {
	max := findMax(array, len(array))

	for i := len(array) - 1; i > 0; i-- {
		utils.Swap(array, max, i)
		max = findMax(array, i)
	}

	return array
}

func findMax(array []int, size int) int {
	var max int

	for i := 1; i < size; i++ {
		if array[i] > array[max] {
			max = i
		}
	}

	return max
}
