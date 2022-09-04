package main

import (
	"fmt"
	"otus-algo/hw06/junior/sort"
	"otus-algo/pkg/utils"
	"time"
)

func main() {
	testSort("Bubble Sort", sort.BubbleSort)
	testSort("Insertion Sort", sort.InsertionSort)
	testSort("Shell Sort", sort.ShellSort)
}

func testSort(name string, sortFn sort.SortFn[int]) {
	sizes := []int{100, 1_000, 10_000, 100_000}

	for _, size := range sizes {
		start := time.Now()

		array := utils.GenerateRandomSample(size)
		sortFn(array)

		elapsed := time.Now().Sub(start).Seconds()
		fmt.Printf("%s sorted array with %d elements, ellapsed %f seconds\n", name, size, elapsed)
	}

	fmt.Println()
}
