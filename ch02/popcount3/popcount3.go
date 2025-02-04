package popcount3

import "fmt"

func PopCount(x uint64) int {
	fmt.Printf("The input: %b\n", x)
	count := 0
	for x != 0 {
		fmt.Printf("Before: %b\n", x)
		x = x & (x - 1) // clear rightmost non-zero bit		
		fmt.Printf(" After: %b\n", x)
		count ++
		fmt.Printf("Count value %v:\n", count)
	}
	return count
}
