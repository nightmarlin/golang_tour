package main

import (
	"fmt"
	"time"
)

func sum(lower, higher int) int {
	sum := 0
	for i := lower; i < higher; i++ {
		sum += i
	}
	return sum
}

func fib(n int) int {
	last2, last, current, count := 0, 0, 1, 1

	for count < n {
		last2, last = last, current
		current = (last + last2)
		count++
	}

	return current
}

func sqrt(val float64, accuracy int) (res float64) {
	res = val / 2

	for i := 0; i < accuracy; i++ {
		res -= ((res * res) - val) / (2 * res)
	}

	return
}

func greeterTime() string {
	t := time.Now()
	g := "Good "
	switch {
	case (5 <= t.Hour() && t.Hour() < 12):
		return g + "Morning"
	case (12 <= t.Hour() && t.Hour() < 17):
		return g + "Afternoon"
	case (17 <= t.Hour() && t.Hour() < 22):
		return g + "Evening"
	default:
		return g + "Night"
	}
}

func deference() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			fmt.Printf("deferred: %d\n", val)
		}(i)
		fmt.Printf("printed: %d\n", i)
	}
}
