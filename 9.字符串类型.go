package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("枫枫知道")
	fmt.Println("'枫枫'知道")
	fmt.Println("\"枫枫\"知道")
	fmt.Println("\\枫枫\\知道")
	fmt.Println("枫枫\n知道")
	fmt.Println("枫枫\t知道")

	var name = `
1
\
2
3
`
	fmt.Println(name)

	fmt.Println(strings.Count("00000000011110000011", "1"))
	fmt.Println(strings.TrimSpace("   zhang san   	"))
	fmt.Println(strings.HasSuffix("xx.jpg", ".jpg"))
	fmt.Println(strings.HasPrefix("user-xxx", "user"))
	fmt.Println(strings.Contains("user-xxx", "x"))
	fmt.Println(strings.Replace("a00a00a00a", "a", "b", 2))
	fmt.Println(strings.ReplaceAll("a00a00a00a", "a", "b"))
	fmt.Println(strings.ToLower("AbB"))
	fmt.Println(strings.ToUpper("AbB"))
	fmt.Println(strings.ToTitle("how are you?"))

}
