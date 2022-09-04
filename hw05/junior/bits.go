package main

func PopCount(mask uint64) int {
	var count int

	for mask > 0 {
		if mask&1 == 1 {
			count++
		}
		mask >>= 1
	}

	return count
}

func PopCount2(mask uint64) int {
	var count int

	for mask > 0 {
		count++
		mask &= mask - 1
	}

	return count
}
