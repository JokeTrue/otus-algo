package main

func isPrime(n int) bool {
	var count int

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}

	return count == 2
}

func PrimeCount(n int) int {
	var count int

	for i := 2; i <= n; i++ {
		if isPrime(i) {
			count++
		}
	}

	return count
}
