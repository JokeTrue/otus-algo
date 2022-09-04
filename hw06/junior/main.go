package main

import (
	"otus-algo/hw06/junior/sort"
	sortPkg "otus-algo/pkg/sort"
)

func main() {
	sortPkg.TestSort("Bubble Sort", sort.BubbleSort)
	sortPkg.TestSort("Insertion Sort", sort.InsertionSort)
	sortPkg.TestSort("Shell Sort", sort.ShellSort)
}
