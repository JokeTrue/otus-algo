package sort

import "otus-algo/pkg/utils"

func HeapSort(array []int) {
	for root := len(array)/2 - 1; root >= 0; root-- {
		heapify(array, root, len(array))
	}

	for i := len(array) - 1; i > 0; i-- {
		utils.Swap(array, 0, i)
		heapify(array, 0, i)
	}
}

func heapify(array []int, root int, size int) {
	left := 2*root + 1
	right := 2*root + 2

	x := root
	if left < size && array[left] > array[x] {
		x = left
	}
	if right < size && array[right] > array[x] {
		x = right
	}
	if x == root {
		return
	}

	utils.Swap(array, x, root)
	heapify(array, x, size)
}
