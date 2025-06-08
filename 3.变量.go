package main

import "fmt"

var age = 26
var (
	x1 = "x"
	x2 = "x"
)

func main() {
	var name string
	name = "枫枫"
	fmt.Println(name)

	var name1 string = "枫枫"
	fmt.Println(name1)

	var name2 = "枫枫"
	fmt.Println(name2)

	name3 := "枫枫"
	fmt.Println(name3)

	fmt.Println(age)

	var a1, a2, a3 = "1", "2", "3"
	fmt.Println(a1, a2, a3)

	fmt.Println(x1, x2)

	var x = "xxx"
	fmt.Println(x)
	x = "xxx"

	const version = "v3"
	//version = "v4"
	fmt.Println(version)

}
