package main

import "fmt"

type user int

var x user

//var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 10
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	y := int(x)
	fmt.Println(y)
	y = 50
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
