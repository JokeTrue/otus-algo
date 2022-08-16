package main

import "fmt"

func CountLucky6NTicket() int {
	var count int

	for a1 := 0; a1 <= 9; a1++ {
		for a2 := 0; a2 <= 9; a2++ {
			for a3 := 0; a3 <= 9; a3++ {
				sumA := a1 + a2 + a3

				for b1 := 0; b1 <= 9; b1++ {
					for b2 := 0; b2 <= 9; b2++ {
						b3 := sumA - b1 - b2

						if b3 >= 0 && b3 <= 9 {
							count++
						}
					}
				}
			}
		}
	}

	return count
}

func main() {
	fmt.Printf("Count of 6N Lucky Tickets is %d\n", CountLucky6NTicket())
}
