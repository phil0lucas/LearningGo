/*
    This is an adaptation from dup2.go to also show the file(s) in which the 
    duplicated string is found.
    1. The program will be simplified to only work with input files.
    2. The file or files containing the duplicates must be stored before 
    printing. Looping through each file in turn means we can use the index
    to determine which file is currently in use. COUNTS store every line
    in every file, only being printed when the count is 2 or more.
    Can we map a string to an array of ints?
*/
package main

import(
	"bufio"
	"fmt"
	"os"
)

func main(){
	counts := make(map[string] int)
	files := os.Args[1:]
	n_files := len(files)
	found := make(map[string] [n_files]string)
	// File processing here
	// The value of the index should show which file is involved
	for index, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		input := bufio.NewScanner(f)
		for input.Scan() {
			counts[input.Text()]++
			found[input.Text()][index] = os.Args[index]
		}		
		f.Close()
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

