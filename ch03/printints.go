package main

import (
	"fmt"
	"bytes"
)

func main() {
	fmt.Println(intsToString([]int{1,2,3,4,5,8,55,-1}))
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		// if it's not the first value
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
