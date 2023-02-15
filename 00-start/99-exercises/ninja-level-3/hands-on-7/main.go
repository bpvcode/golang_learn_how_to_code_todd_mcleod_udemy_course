package main

import "fmt"

func main() {
	i := "NOT Hello"
	if i == "Hello" {
		fmt.Println("IF")
	} else if i == "NOT Hello" {
		fmt.Println("ELSE IF")
	} else {
		fmt.Println("ELSE")
	}
}
