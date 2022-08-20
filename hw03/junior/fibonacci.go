package main

func Fibonacci(n int) int {
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		temp := a
		a = b
		b = temp + a
	}

	return a
}

func FibonacciRecursive(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return FibonacciRecursive(n-2) + FibonacciRecursive(n-1)
}
