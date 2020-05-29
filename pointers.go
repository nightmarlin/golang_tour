package main

import "fmt"

func pointerisms(i int) {

	fmt.Printf("The value stored at %p is (%T) %v\n", &i, i, i)
}
