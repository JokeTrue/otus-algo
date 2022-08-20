package main

import "fmt"

func main() {
	fmt.Printf("Fibonacci %dth is %d === 143\n", 12, FibonacciBinet(12))
	fmt.Printf("Fibonacci %dth is %d === 4181\n", 19, FibonacciBinet(19))
	fmt.Printf("Fibonacci %dth is %d === 43367\n\n", 24, FibonacciBinet(24))

	fmt.Printf("Count of Primes less then %d is %d === 4\n", 10, Eratosthenes(10))
	fmt.Printf("Count of Primes less then %d is %d === 25\n\n", 100, Eratosthenes(100))
}
