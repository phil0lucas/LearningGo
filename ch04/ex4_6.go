// Compress excess white space in UTF-8 byte slice


package main

import (
	"fmt"
	"bytes"
)

func main() {
	// This is a byte slice
	x := []byte("aaa       aaa bbb    cc  bbb")
	//fmt.Println(bytes.Trim(x, "\x00"))
	fmt.Printf("Changed slice: %s\n",compress(x))
}

func compress(s []byte) []byte{
	var buf bytes.Buffer
	fmt.Printf("In function: %s\n",s)
	ii := 0
	var last byte
	for _, v := range s {
		fmt.Println(v)
		if (string(v) == "") & (v == last) {
		} else {
			buf.WriteString(v)
		}
		last = v
		ii++
	}
	return string(buf)
}
