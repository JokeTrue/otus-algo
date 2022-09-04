package sort

import "otus-algo/pkg/utils"

func ShellSort(array []int) {
	for gap := len(array) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(array); i++ {
			for j := i; j >= gap && array[j-gap] > array[j]; j -= gap {
				utils.Swap(array, j, j-gap)
			}
		}
	}
}
