package sort

func QuickSort(array []int) []int {
	return quickSort(array, 0, len(array)-1)
}

func quickSort(array []int, low int, high int) []int {
	var p int

	if low < high {
		array, p = partition(array, low, high)
		array = quickSort(array, low, p-1)
		array = quickSort(array, p+1, high)
	}

	return array
}

func partition(array []int, low int, high int) ([]int, int) {
	i := low
	pivot := array[high]

	for j := low; j < high; j++ {
		if array[j] < pivot {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}
	array[i], array[high] = array[high], array[i]

	return array, i
}
