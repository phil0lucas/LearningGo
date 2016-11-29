package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/phil0lucas/LearningGo/tempconv2"
)

func main () {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv2.Fahrenheit(t)
		c := tempconv2.Celsius(t)
		k := tempconv2.Kelvin(t)
		fmt.Printf(
			"%s = %s = %s %s = %s = %s %s = %s = %s\n",
			f, tempconv2.FtoC(f), tempconv2.FtoK(f),
			c, tempconv2.CtoF(c), tempconv2.CtoK(c),
			k, tempconv2.KtoF(k), tempconv2.KtoC(k))
	}
}
