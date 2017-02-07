package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func main(){
	fmt.Println(os.Args[1:2])					// [123456]
	fmt.Println(len(os.Args[1:2]))				// 1
	fmt.Println(reflect.TypeOf(os.Args[1:2])) 	// []string
	s := strconv.Itoa(os.Args[1:2])
	//t := s[:]
	fmt.Println(reflect.TypeOf(s))
 	fmt.Println(comma(s))
}	

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}