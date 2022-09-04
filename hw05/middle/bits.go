package main

func PopCount2(mask uint64) int {
	var count int

	for mask > 0 {
		count++
		mask &= mask - 1
	}

	return count
}

func PopCount3(mask uint64) int {
	var count int

	for mask > 0 {
		count += bits[mask&255]
		mask >>= 8
	}

	return count
}
