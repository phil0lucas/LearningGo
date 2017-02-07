// echo prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for ii := 1; ii < len(os.Args); ii++ {
		s += sep + os.Args[ii]
		sep = " "
	}
	fmt.Println(s)
}
