// Slice of strings - eliminate (in place) adjacent duplicates


package main

import "fmt"

func main() {
	// This is a slice not an array
	x := []string{"aaa", "aaa", "bbb", "ccc", "bbb", "aaa", "aaa", "aaa"}
	fmt.Printf("Changed slice: %s\n",nodup(x))
}

func nodup(s []string) []string {
	fmt.Printf("In function: %s\n",s)
	j := 0
	last := ""
	for _, v := range s {
		if v != last{
			s[j] = v
			last = v
			j++
		}
	}
	return s[:j]
}