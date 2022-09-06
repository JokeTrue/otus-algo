package sort

import (
	"fmt"
	"otus-algo/pkg/utils"
	"time"
)

func TestSort(name string, sortFn SortFn[int]) {
	sizes := []int{100, 1_000, 10_000, 100_000, 1_000_000}

	for _, size := range sizes {
		start := time.Now()

		array := utils.GenerateRandomSample(size)
		array = sortFn(array)

		elapsed := time.Now().Sub(start).Seconds()
		fmt.Printf("%s sorted array with %d elements, ellapsed %f seconds\n", name, size, elapsed)
	}

	fmt.Println()
}
