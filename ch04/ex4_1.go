package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("xx"))
	c2 := sha256.Sum256([]byte("xx"))
	fmt.Printf("%8b\n%8b\n%t\n%T\n", c1, c2, c1==c2, c1)
	
	// So, c1 and c2 are arrays of 32 * 8 bit unsigned bytes = 256 bits
	count := 0
	for i := 0; i < 32; i++ {
		c1Byte := c1[i]
		c2Byte := c2[i]
		for j := 0; j < 8; j++ {
			mask := byte(1 << uint(j))
			if (c1Byte & mask) != (c2Byte & mask) {
				count++
			}
		}
	}
	fmt.Printf("%v\n", count)
}