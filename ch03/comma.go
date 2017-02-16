package main

import (
	"fmt"
	"os"
	"bytes"
)

func main() {
	arg := os.Args[1]
	fmt.Printf("Argument value %s \n", arg)
	var s string
	s = comma(arg)
	fmt.Printf("Arg value from basename1 %s \n", s)
}


func comma(s string) string {
	l := len(s)
	if l > 3 {
		b := bytes.NewBufferString(s)
		for i := l-1; i = 0; i--{
			token := b[i]
			fmt.Printf("Token %s \n", token)
		}
	}
	return s
}
