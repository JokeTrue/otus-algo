package main

import "fmt"

func main() {
	fmt.Printf("%d^%d is %d === 1\n", 2, 0, Power(2, 0))
	fmt.Printf("%d^%d is %d === 256\n\n", 2, 8, Power(2, 8))

	fmt.Printf("%dth number of Fibonacci is %d === 8\n", 6, Fibonacci(6))
	fmt.Printf("%dth number of Fibonacci is %d === 144\n", 12, Fibonacci(12))
	fmt.Printf("%dth number of Fibonacci is %d === 46368\n\n", 24, Fibonacci(24))

	fmt.Printf("[recursive] %dth number of Fibonacci is %d === 8\n", 6, FibonacciRecursive(6))
	fmt.Printf("[recursive] %dth number of Fibonacci is %d === 144\n", 12, FibonacciRecursive(12))
	fmt.Printf("[recursive] %dth number of Fibonacci is %d === 46368\n\n", 24, FibonacciRecursive(24))

	fmt.Printf("PrimeCount of %d is %d === 4\n", 10, PrimeCount(10))
	fmt.Printf("PrimeCount of %d is %d === 25\n\n", 100, PrimeCount(100))
}
