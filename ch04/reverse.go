package main

import "fmt"

func main() {
	x:= [...]int{0, 1, 2, 3, 4, 5, 99}
	reverse(x[:])
	fmt.Println(x)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i<j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}