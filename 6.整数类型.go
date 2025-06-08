package main

import "fmt"

func main() {
	var a uint8 = 255
	// 0 0 0 0 0 0 0 0  = 0
	// 1 1 1 1 1 1 1 1  = 255
	fmt.Printf("%0b %d\n", a, a)

	var b int8 = 127
	// 0    0 0 0 0 0 0 0   = 0
	// 0    1 1 1 1 1 1 1   = 127
	// 1    0 0 0 0 0 0 0   = -128
	fmt.Printf("%0b %d\n", b, b)
	b = -128
	fmt.Printf("%0b %d\n", b, b)

	var c = 9223372036854775807
	fmt.Println(c)
}
