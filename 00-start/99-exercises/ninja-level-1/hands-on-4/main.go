package main

import "fmt"

type user int

var x user

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 10
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
