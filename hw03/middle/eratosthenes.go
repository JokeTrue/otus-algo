package main

func Eratosthenes(n int) int {
	prime := make([]bool, n+1)
	for i := 0; i < n+1; i++ {
		prime[i] = false
	}

	for i := 2; i*i <= n; i++ {
		if prime[i] == false {
			for j := i * 2; j <= n; j += i {
				prime[j] = true
			}
		}
	}

	var count int
	for i := 2; i <= n; i++ {
		if prime[i] == false {
			count++
		}
	}

	return count
}
