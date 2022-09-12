package main

import (
	"otus-algo/hw09/junior/sort"
	sortPkg "otus-algo/pkg/sort"
)

func main() {
	sortPkg.TestSort("Counting Sort", sort.CountingSort)
	sortPkg.TestSort("Bucket Sort", sort.BucketSort)
	sortPkg.TestSort("Radix Sort", sort.RadixSort)
}
