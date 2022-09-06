package sort

import "otus-algo/pkg/utils"

func InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := i - 1; j >= 0 && array[j] > array[j+1]; j-- {
			utils.Swap(array, j, j+1)
		}
	}

	return array
}
