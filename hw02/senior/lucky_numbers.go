package main

func generateArray(arr []int) []int {
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
		array = generateArray(array)
	}

	var result int
	for _, value := range array {
		result += value * value
	}

	return result
}
