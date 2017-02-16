package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1]
	fmt.Printf("Argument value %s \n", arg)
	var s string
	s = basename2(arg)
	fmt.Printf("Arg value from basename1 %s \n", s)
}


func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	dot := strings.LastIndex(s, ".")
	s = s[:dot]
	return s
}
