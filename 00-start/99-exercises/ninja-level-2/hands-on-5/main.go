package main

import "fmt"

func main() {
	a := "Nice"
	x := `Hello this is a raw string literal -> ` + a

	fmt.Println(x)
}
