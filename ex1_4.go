/*
   This is an adaptation from dup2.go to also show the file(s) in which the
   duplicated string is found.
   1. The program will be simplified to only work with input files.
   2. The file or files containing the duplicates must be stored before
   printing. Looping through each file in turn means we can use the index
   to determine which file is currently in use.
   COUNTS: stores every unique line in every input file -> counts of that line
   only being printed when the count is 2 or more.
   COUNTSFILES maps each unique line to the files it appears in.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	countsFiles := make(map[string][]string)
	files := os.Args[1:]

	for _, arg := range files {
		name := arg
		fmt.Println("%s", arg)
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		input := bufio.NewScanner(f)
		for input.Scan() {

			text := input.Text()
			counts[text]++
			if !stringInSlice(name, countsFiles[text]) {
				countsFiles[text] = append(countsFiles[text], name)
			}
		}
		f.Close()
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, countsFiles[line])
		}
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
