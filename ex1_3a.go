// echo prints its command-line arguments
package main

import (
	"os" 
	"fmt"
	"time"
)

func main(){
    start := time.Now()
	var s, sep string
	for ii := 1; ii < len(os.Args); ii++ {
		s += sep + os.Args[ii]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(time.Since(start).Seconds())
}
