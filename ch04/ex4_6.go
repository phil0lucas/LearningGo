// Compress excess white space in UTF-8 byte slice


package main

import (
	"fmt"
	"bytes"
	"unicode"
)

func main() {
	// This is a byte slice
	x := []byte("aaa       aaa bbb    cc  bbb")
	fmt.Printf("Original slice: %s\n", x)
	fmt.Printf("Changed slice: %s\n",compress(x))
}

func compress(s []byte) string{
	var buf bytes.Buffer
	//fmt.Printf("In function: %s\n",s)
	ii := 0
	var last byte
	for _, v := range s {
		// fmt.Println(v)
		if (unicode.IsSpace(rune(v))) && (v == last) {
		} else {
			buf.WriteByte(v)
		}
		last = v
		ii++
	}
	return buf.String()
}
