package main

func GetKingBitboardMoves(pos int) uint64 {
	k := uint64(1 << pos)

	noA := uint64(0xfefefefefefefefe)
	noH := uint64(0x7f7f7f7f7f7f7f7f)

	kA := k & noA
	kH := k & noH

	mask := (kA << 7) | (k << 8) | (kH << 9)
	mask |= (kA >> 1) /*      */ | (kH << 1)
	mask |= (kA >> 9) | (k >> 8) | (kH >> 7)

	return mask
}

func GetKnightBitboardMoves(pos int) uint64 {
	k := uint64(1 << pos)

	noA := uint64(0xfefefefefefefefe)
	noAB := uint64(0xfcfcfcfcfcfcfcfc)

	noH := uint64(0x7f7f7f7f7f7f7f7f)
	noGH := uint64(0x3f3f3f3f3f3f3f3f)

	mask := noGH & (k<<6 | k>>10)
	mask |= noH & (k<<15 | k>>17)
	mask |= noA & (k<<17 | k>>15)
	mask |= noAB & (k<<10 | k>>6)

	return mask
}
