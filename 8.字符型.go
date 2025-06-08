package main

import "fmt"

func main() {
	var c byte
	c = 'a'
	fmt.Println(c)
	fmt.Printf("%c\n", c)
	c = 98
	fmt.Println(c)
	fmt.Printf("%c\n", c)

	var b rune
	b = '枫'
	fmt.Println(b)
	fmt.Printf("%c\n", b)
	b = '知'
	fmt.Println(b)
	fmt.Printf("%c\n", b)

}
