package main

import (
	"fmt"
	"github.com/phil0lucas/LearningGo/popcount"
	"os"
	"strconv"
)

func main() {
	u, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex2_3: %v\n", err)
	}
	fmt.Println(popcount.PopCount(u))
}
