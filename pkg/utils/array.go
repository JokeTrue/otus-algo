package utils

import (
	"math/rand"
	"time"
)

func Swap[T any](array []T, i int, j int) {
	tmp := array[i]
	array[i] = array[j]
	array[j] = tmp
}

func GenerateRandomSample(size int) []int {
	rand.Seed(time.Now().Unix())

	sample := make([]int, 0, size)
	for i := 0; i < size; i++ {
		sample = append(sample, rand.Intn(999))
	}

	return sample
}
