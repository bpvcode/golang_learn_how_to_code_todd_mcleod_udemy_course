package main

import "fmt"

var x = 42

const y = "James Bond"

var z = true

func main() {
	s := fmt.Sprintf("%v\t %v\t %v\t", x, y, z)
	fmt.Println(s)
	fmt.Printf("%T\n", s)
}
