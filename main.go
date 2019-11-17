package main

import (
	"fmt"
	"log"

	"github.com/jbrukh/go-banzhaf"
)

func main() {
	weights := make([]uint64, 20000)
	balance := uint64(2000)
	for i := range weights {
		weights[i] = balance
	}
	quota := uint64(40000000)/2 + 1

	bi, ok := banzhaf.Banzhaf(weights, quota, true)
	if !ok {
		log.Fatal("error in Banzhaf")
	}
	for i, b := range bi {
		fmt.Printf("%d, %.2f\n", i, b)
	}
}
