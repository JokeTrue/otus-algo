package sort

import "otus-algo/hw06/junior/sort"

const BucketSize = 5

func BucketSort(array []int) []int {
	var (
		min int
		max int
	)
	for _, value := range array {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	buckets := make([][]int, (max-min)/BucketSize+1)
	for i := 0; i < len(buckets); i++ {
		buckets[i] = make([]int, 0)
	}

	for _, value := range array {
		idx := (value - min) / BucketSize
		buckets[idx] = append(buckets[idx], value)
	}

	newArray := make([]int, 0, len(array))
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			bucket = sort.InsertionSort(bucket)
			newArray = append(newArray, bucket...)
		}
	}

	return newArray
}
