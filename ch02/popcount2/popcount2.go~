package popcount2
import "fmt"
func PopCount(x uint64) int {
	count := 0
	var i int8
	for i = 0; i < 64; i++ {
		fmt.Printf("Before shift: %b\n", x)
		if hasBit(x, 0) {
			count ++
		}
		x = x >> 1		
		fmt.Printf(" After shift: %b\n", x)
	}
	return count
}

func hasBit(n uint64, pos uint) bool {
	fmt.Printf(" Input value: %b\n", n)
	fmt.Printf(" Shift compn: %b\n", (1 << pos))
    val := n & (1 << pos)
    fmt.Printf(" After AND  : %b\n", val)
    return (val > 0)
}


