package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	a := 1.0
	fmt.Println(a)
	b := float64(1)
	fmt.Println(b)

	c := 1
	d := 1.0
	e := float64(c) + d
	fmt.Println(e)

	fmt.Println(math.Inf(1) > 10000000000000000.1)
	fmt.Println(math.Inf(-1) < -10000000000000.1)
	fmt.Println(math.IsInf(math.Inf(1), 1))

	fmt.Println(math.NaN())
	fmt.Println(math.IsNaN(1.1))

}
