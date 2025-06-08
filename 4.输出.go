package main

import "fmt"

func main() {
	var name = "dx"
	fmt.Println(name)
	fmt.Println("xxx")
	fmt.Println(123)

	fmt.Print("1", "xxx")
	fmt.Print("2")

	fmt.Printf("\"枫枫\"")
	name = "枫枫"
	fmt.Printf("%q %T\n", name, name)

	var age = 12
	fmt.Printf("%s\n", age)

	s := fmt.Sprintf("我的年龄是%d", age)
	fmt.Println(s)

}
