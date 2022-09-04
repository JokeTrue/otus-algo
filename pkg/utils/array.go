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
	return rand.Perm(size)
}
