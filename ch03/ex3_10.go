package main

import (
	"fmt"
	"bytes"
//	"strconv"
)

func main() {
	fmt.Println(comma("123"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))	
}

func comma(s string) string {
	var buf bytes.Buffer
	
	counter := 0
	for i := len(s)-1; i >= 0; i-- {
		//fmt.Println("Current Value " + s[i:i+1])
		//fmt.Println("Current Index " + strconv.Itoa(i))
		counter++
		fmt.Fprintf(&buf, "%s", s[i:i+1])
		if counter % 3 == 0 && counter < len(s) {
			buf.WriteString(",")
		}
	}
	return reverse(buf.String())
}

func reverse(s string) string {
	var result string
  	for _,v := range s {
    	result = string(v) + result
  	}
    return result
}
