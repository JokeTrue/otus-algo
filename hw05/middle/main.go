package main

import "fmt"

var bits []int

func init() {
	for b := 0; b < 256; b++ {
		bits = append(bits, PopCount2(uint64(b)))
	}
}

func main() {
	fmt.Println(GetRookBitboardMoves(58))

	fmt.Println(PopCount2(uint64(299689933283329)))
	fmt.Println(PopCount3(uint64(299689933283329)))
}
