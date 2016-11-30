package main

import (
	"fmt"
	"os"
	"strconv"	
)

type Pound float64
type Kilogram float64

const (
	lb2kg Pound = 0.45359237
	kg2lb Kilogram = 2.20462
)

func convlb2kg (p Pound) Kilogram {
	return Kilogram(p * lb2kg)
}

func convkg2lb (k Kilogram) Pound {
	return Pound(k * kg2lb)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%6.2fkG", k)
}

func (p Pound) String() string {
	return fmt.Sprintf("%6.2flb", p)
}

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		p := Pound(t)
		k := Kilogram(t)
		
		fmt.Printf(
			"%s = %s,  %s = %s\n",
			p, convlb2kg(p), k, convkg2lb(k))
	}
}
