package popcount1

var pc [256]byte

func init() {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	count := 0
	var i uint8
	for i = 0; i < 7; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}
