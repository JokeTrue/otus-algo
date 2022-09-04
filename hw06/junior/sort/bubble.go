package sort

import "otus-algo/pkg/utils"

func BubbleSort(array []int) {
	for i := 1; i < len(array); i++ {
		for j := 0; j < len(array)-i; j++ {
			if array[j] > array[j+1] {
				utils.Swap(array, j, j+1)
			}
		}
	}
}
