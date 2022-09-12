package sort

import "otus-algo/pkg/utils"

func RadixSort(array []int) []int {
	digit := 1
	size := len(array)

	max := utils.FindMax(array)
	almostSortedArray := make([]int, size, size)

	for max/digit > 0 {
		bucket := [10]int{0}

		for i := 0; i < size; i++ {
			bucket[(array[i]/digit)%10]++
		}

		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i-1]
		}

		for i := size - 1; i >= 0; i-- {
			bucket[(array[i]/digit)%10]--
			almostSortedArray[bucket[(array[i]/digit)%10]] = array[i]
		}

		for i := 0; i < size; i++ {
			array[i] = almostSortedArray[i]
		}

		digit *= 10
	}

	return array
}
