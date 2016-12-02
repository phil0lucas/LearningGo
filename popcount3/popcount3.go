package popcount3

func PopCount(x uint64) int {
	count := 0
	var i int8
	for i = 0; i < 64; i++ {
		for x != 0 {
			x = x & (x - 1) // clear rightmost non-zero bit		
			count ++
		}
	}
	return count
}
