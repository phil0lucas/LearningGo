package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma (s string) string {
	t, pref := signPrefix(s)
	fmt.Println(t)
	fmt.Println(pref)
	t2, dec := afterDec(t)
	fmt.Println(t2)
	fmt.Println(dec)
	withCommas := body(t2)
	fmt.Println(withCommas)	
	return pref + withCommas + dec
}

func signPrefix(s string) (string, string) {
	if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "+"){
		fmt.Println("Sign prefix seen")
		return s[1:], s[0:1]
	}
	return s, ""
}

func afterDec(s string) (string, string) {
	if strings.Contains(s, ".") {
		fmt.Println("Decimal seen")
		index := strings.Index(s, ".")
		return s[0:index], s[index:]
	}
	return s, ""
}


// This enhanced version of Comma is to handle optional leading sign and 
// decimal places
func body(s string) string {	
	n := len(s)
	if n <= 3 {
		return s
	}
	return body(s[:n-3]) + "," + s[n-3:]
}
