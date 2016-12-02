package main

import (
	"fmt"
	"github.com/phil0lucas/LearningGo/popcount3"
	"os"
	"strconv"
)

func main() {
	u, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex2_5: %v\n", err)
	}
	fmt.Println(popcount3.PopCount(u))
}
