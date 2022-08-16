package main

import "fmt"

func GenerateArray(arr []int) []int {
	newLength := len(arr) + 9
	newArray := make([]int, 0, newLength)

	for i := 0; i < newLength; i++ {
		var value int

		for j := 0; j < 10; j++ {
			index := i - j
			if index >= 0 && len(arr) > index {
				value += arr[index]
			}
		}

		newArray = append(newArray, value)
	}

	return newArray
}

func CountLuckyTickets(num int) int {
	array := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		array = append(array, 1)
	}

	for i := 0; i < num/2-1; i++ {
		array = GenerateArray(array)
	}

	var result int
	for _, value := range array {
		result += value * value
	}

	return result
}

func main() {
	fmt.Printf("Count of %dN Lucky Tickets is %d\n", 2, CountLuckyTickets(2))
	fmt.Printf("Count of %dN Lucky Tickets is %d\n", 4, CountLuckyTickets(4))
	fmt.Printf("Count of %dN Lucky Tickets is %d\n", 6, CountLuckyTickets(6))
	fmt.Printf("Count of %dN Lucky Tickets is %d\n", 8, CountLuckyTickets(8))
	fmt.Printf("Count of %dN Lucky Tickets is %d\n", 10, CountLuckyTickets(10))
	fmt.Printf("Count of %dN Lucky Tickets is %d\n", 12, CountLuckyTickets(12))
	fmt.Printf("Count of %dN Lucky Tickets is %d\n", 14, CountLuckyTickets(14))
	fmt.Printf("Count of %dN Lucky Tickets is %d\n", 16, CountLuckyTickets(16))
	fmt.Printf("Count of %dN Lucky Tickets is %d\n", 18, CountLuckyTickets(18))
	fmt.Printf("Count of %dN Lucky Tickets is %d\n", 20, CountLuckyTickets(20))
}
