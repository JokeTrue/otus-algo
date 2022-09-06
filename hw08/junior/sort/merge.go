package sort

func MergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	mid := len(array) / 2
	return merge(MergeSort(array[:mid]), MergeSort(array[mid:]))
}

func merge(left []int, right []int) []int {
	var (
		i int
		j int
	)

	size := len(left) + len(right)
	newArray := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			newArray[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			newArray[k] = left[i]
			i++
		} else if left[i] < right[j] {
			newArray[k] = left[i]
			i++
		} else {
			newArray[k] = right[j]
			j++
		}
	}

	return newArray
}
