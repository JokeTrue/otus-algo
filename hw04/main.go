package main

import (
	"fmt"
	arrays2 "otus-algo/hw04/arrays"
	"time"
)

func main() {
	single := arrays2.NewSingleArray[int]()
	testArrayAdd("Single Array", single, 1_000)
	testArrayAdd("Single Array", single, 10_000)
	testArrayAdd("Single Array", single, 100_000)
	testArrayAdd("Single Array", single, 1_000_000)
	fmt.Println()

	vector := arrays2.NewVectorArray[int](1000)
	testArrayAdd("Vector Array", vector, 1_000)
	testArrayAdd("Vector Array", vector, 10_000)
	testArrayAdd("Vector Array", vector, 100_000)
	testArrayAdd("Vector Array", vector, 1_000_000)
	testArrayAdd("Vector Array", vector, 10_000_000)
	fmt.Println()

	factor := arrays2.NewFactorArray[int](2)
	testArrayAdd("Factor Array", factor, 1_000)
	testArrayAdd("Factor Array", factor, 10_000)
	testArrayAdd("Factor Array", factor, 100_000)
	testArrayAdd("Factor Array", factor, 1_000_000)
	testArrayAdd("Factor Array", factor, 10_000_000)
	fmt.Println()
}

func testArrayAdd(name string, array arrays2.Array[int], count int) {
	start := time.Now()

	for i := 0; i < count; i++ {
		array.Add(i)
	}

	elapsed := time.Now().Sub(start).Seconds()
	fmt.Printf("%s: added %d items, ellapsed %f seconds\n", name, count, elapsed)
}
