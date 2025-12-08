package bit_manip


func reverseBits(x uint32) uint32 {
	var res uint32
	for range 32 {
		res <<= 1
		res |= (x & 1)
		x >>= 1
	}
	return res
}

func reverseBits2(x uint32) uint32 {
	var res uint32
	for range 32 {
		res = (res << 1) | (x & 1)
		x >>= 1
	}
	return res
}
