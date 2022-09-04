package main

import "fmt"

func main() {
	fmt.Println(GetKingBitboardMoves(31))
	fmt.Println(GetKnightBitboardMoves(45))

	fmt.Println(PopCount(uint64(299689933283329)))
	fmt.Println(PopCount2(uint64(299689933283329)))
}
