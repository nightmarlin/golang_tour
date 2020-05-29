package main

import (
	"fmt"
	"math"
	"math/rand"
)

func printRandomInt(maxVal int) {
	fmt.Println(rand.Intn(maxVal), "is a random number")
}

func printSqrt(val float64) {
	fmt.Printf("And %g is the sqrt of %g", math.Sqrt(val), val)
}
