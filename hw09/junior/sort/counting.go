package sort

import "otus-algo/pkg/utils"

func CountingSort(array []int) []int {
	max := utils.FindMax(array)
	counter := make([]int, max+1)

	for _, value := range array {
		counter[value]++
	}

	for i := 1; i < len(counter); i++ {
		counter[i] += counter[i-1]
	}

	newArray := make([]int, len(array)+1)
	for i := 0; i < len(array); i++ {
		newArray[counter[array[i]]] = array[i]
		counter[array[i]] = counter[array[i]] - 1
	}

	return newArray
}
