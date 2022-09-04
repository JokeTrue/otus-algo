package main

func Fibonacci(n int) int {
	first := 0
	second := 1

	if n == 0 {
		return first
	} else if n <= 2 {
		return second
	}

	for i := 2; i <= n; i++ {
		second = second + first
		first = second - first
	}

	return second
}

func FibonacciRecursive(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return FibonacciRecursive(n-2) + FibonacciRecursive(n-1)
}
