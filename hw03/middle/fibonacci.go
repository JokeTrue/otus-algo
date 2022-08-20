package main

import "math"

func FibonacciBinet(n int) int {
	return int((math.Pow(math.Phi, float64(n)) - math.Pow(1-math.Phi, float64(n))) / math.Sqrt(5))
}
