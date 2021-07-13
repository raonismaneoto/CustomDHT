package main

import (
	"fmt"
	"math"
)

type Tst struct {
	id string
}

func main() {
	fmt.Println((1 + int64(math.Pow(2, float64(1-1))))%int64(math.Pow(2, float64(32))))
}
