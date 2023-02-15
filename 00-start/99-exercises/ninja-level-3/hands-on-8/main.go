package main

import "fmt"

func main() {
	i := "TEST"
	switch {
	case i == "Hello":
		fmt.Println("Hello")
	case i == "Bye":
		fmt.Println("Bye")
	case i == "Xau":
		fmt.Println("Xau")
	case i == "Ola":
		fmt.Println("Ola")
	default:
		fmt.Println("NONE")
	}
}
