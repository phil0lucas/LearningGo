// Charcount computes counts of Unicode characters.
package main

import (
	"log"
	"os"
	"fmt"
	"bufio"
)

func main() {
	// Open file and create scanner on top of it
    file, err := os.Open("Swanns_Way_English.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
	
    // Default scanner is bufio.ScanLines. Lets use ScanWords.
    // Could also use a custom function of SplitFunc type
    scanner.Split(bufio.ScanWords)	
	
	// Count the words
	/*
	count := 0
	for scanner.Scan(){
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)
	*/
	
	// Count each different word
	wmap := make(map[string]int)
	for scanner.Scan(){
		word := scanner.Text()
		//fmt.Println(word)
		wmap[word]++
	}
	fmt.Println(wmap)
	
}
