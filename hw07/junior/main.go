package main

import (
	"otus-algo/hw07/junior/sort"
	sortPkg "otus-algo/pkg/sort"
)

func main() {
	sortPkg.TestSort("Selection Sort", sort.SelectionSort)
	sortPkg.TestSort("Heap Sort", sort.HeapSort)
}
