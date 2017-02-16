// echo3 prints its command-line arguments
// this version takes advantage of strings.Join
package main

import (
	"os" 
	"fmt"
	"strings"
	"time"
)

func main(){
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Since(start).Seconds())
}
