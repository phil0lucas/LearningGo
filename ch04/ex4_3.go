package main

import "fmt"

func main() {
	x := [...]int{0, 1, 2, 3, 4, 5, 99}
	fmt.Printf("Type: %T\n",x)
	reverse(&x)
	fmt.Printf("Changed array: %v\n",x)
}

func reverse(s *[7]int) {
	for i, j := 0, len(s)-1; i<j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}