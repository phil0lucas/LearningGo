package main

import(
	"fmt"
)

func main() {
	x := 1
	p := &x
	fmt.Printf("Directly: %d\n ", x)
	fmt.Printf("Indirectly: %d\n ", *p)
	*p = 99
	fmt.Printf("Update indirectly %d\n ", *p)
	
	v := 1
	incr(&v)
	fmt.Println(incr(&v))
}

func incr(p *int) int {
	*p++
	return *p
}

