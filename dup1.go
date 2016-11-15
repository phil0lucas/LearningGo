/*
 os.Stdin is the terminal input.
 Run the program - which pauses for input - and at the terminal type something like:
aaa<enter>
aaa<enter>
bbb<enter>
ccc<enter>
aaa<enter>
bbb<enter>
What they don't tell you is to end the stream from os.Stdin by Control-D
*/
package main

import(
	"bufio"
	"fmt"
	"os"
)

func main(){
	// Create a string to int mapping (ike a value format in SAS terms)
	counts := make(map[string] int)
	// This captures the os.Stdin input
	input := bufio.NewScanner(os.Stdin)
	// each call the Scan() reads a new line
	// input.Text() is the text in the line
	// the text then becomes part of the string side of the map. The default
	// integer associated with it is 0, so ++ increments. Hence if a value > 1 
	// is seen in the int side of the map, it must be a duplicate
	for input.Scan(){
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
