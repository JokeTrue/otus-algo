package main

import (
	"otus-algo/hw08/junior/sort"
	sortPkg "otus-algo/pkg/sort"
)

func main() {
	sortPkg.TestSort("Merge Sort", sort.MergeSort)
	sortPkg.TestSort("Quick Sort", sort.QuickSort)
}
