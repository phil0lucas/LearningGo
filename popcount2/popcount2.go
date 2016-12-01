package popcount2

func PopCount(x uint64) int {
	count := 0
	var i int8
	for i = 0; i < 64; i++ {
		if hasBit(x, 0) {
			count ++
		}
		x = x >> 1		
	}
	return count
}

func hasBit(n uint64, pos uint) bool {
    val := n & (1 << pos)
    return (val > 0)
}


