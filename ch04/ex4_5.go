// Slice of strings - eliminate (in place) adjacent duplicates


package main

import "fmt"

func main() {
	x := []string{"aaa", "aaa", "bbb", "ccc", "bbb", "aaa"}
	//fmt.Printf("Type: %T\n",x)
	fmt.Println(dedup(x))
	//nodup(&x)
	//fmt.Printf("Changed slice: %s\n",x)
}
/*
func nodup(s *[6]string){
	fmt.Printf("In function: %s\n",s)
	for i = 1; i < len(s); i++ {
		if s[i] == s[i - 1]{
		}
	}
}
*/

func dedup(items []string) []string {
	i := 0
	last := ""; 
	for _, s := range items {
		if s != last {
			items[i] = s
			last = s
			i++
		}
	}
	return items[:i]
}