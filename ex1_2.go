// echo2 prints its command-line arguments
// ex1_2.go print the index and args one per line
// Note index start at 0, even though the slice is from the second token
package main

import (
	"os" 
	"fmt"
)

func main(){
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}
}
