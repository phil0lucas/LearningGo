// Not sure this is good code


package main

import "fmt"

func main() {
	x := []int{0, 1, 2, 3, 4, 5, 99}
	fmt.Printf("Type: %T\n",x)
	rotSlice := rotate(x, 0)
	fmt.Printf("Changed slice: %v\n",rotSlice)
}

func rotate(s []int, npos int) []int {
	lenSlice := len(s)
	newSlice := make([]int, 0)
	offset := lenSlice - npos
	front := s[offset:]
	back := s[:offset]
	
	for i := 0; i < len(front); i++ {
		var x int
		x = front[i]
		newSlice = append(newSlice, x)
	}
	
	for i := 0; i < len(back); i++ {
		var x int
		x = back[i]
		newSlice = append(newSlice, x)
	}	
		
	//fmt.Printf("Input slice length %v\n", lenSlice)
	
	//fmt.Printf("Front %v, Back %v", s[offset:], s[:offset])
	
	//fmt.Printf("Output slice %v\n", newSlice)
	
	//copy(newSlice, s[offset:])
	//fmt.Printf("After copy  %v\n", newSlice)
	
	//for i := 0; i < len(s[:offset]); i++ {
	//	var x int
	//	x = s[:offset][i]
	//	fmt.Printf("Bits to append %v\n", x)
	//	newSlice = append(newSlice, x)
	//}
	return newSlice
}