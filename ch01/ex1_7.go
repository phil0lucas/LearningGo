package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main(){
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		
		_, err2 := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		
		if err2 != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err2)
			os.Exit(1)
		}
		//fmt.Printf("%s", b)
	}
}
