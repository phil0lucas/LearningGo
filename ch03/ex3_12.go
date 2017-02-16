package main

import (
	"fmt"
	"os"
	"strings"
	"sort"
)

func main() {
	primaryStr := os.Args[1]
	fmt.Println(primaryStr)
	secondaryStr := os.Args[2]
	fmt.Println(secondaryStr)	
	
	if len(primaryStr) != len(secondaryStr) {
		fmt.Println("Input Strings are of different lengths and thus not anagrams")
	} else {
		aLetters := strings.Split(primaryStr, "")
		sort.Strings(aLetters)
		fmt.Printf("%q\n", aLetters)
	
		bLetters := strings.Split(secondaryStr, "")
		sort.Strings(bLetters)
		fmt.Printf("%q\n", bLetters)
	
		for i := 0; i < len(aLetters); i++ {
			if aLetters[i] != bLetters[i]{
				fmt.Println("The input strings are not anagrams of each other")
			}
		}
	}
}
