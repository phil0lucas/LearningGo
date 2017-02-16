package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1]
	fmt.Printf("Argument value %s \n", arg)
	var s string
	s = basename1(arg)
	fmt.Printf("Arg value from basename1 %s \n", s)
	//fmt.Printf("Base name without directory or extension: %s ", basename1(arg))
}


func basename1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
		}
	}
	
	
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
		}
	}
	
	return s
}
