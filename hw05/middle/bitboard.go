package main

func GetRookBitboardMoves(pos int) uint64 {
	y := uint64(pos / 8)
	x := uint64(pos % 8)

	hLine := uint64(0xff)
	vLine := uint64(0x101010101010101)

	return (vLine << (1 * x)) ^ (hLine << (8 * y))
}
