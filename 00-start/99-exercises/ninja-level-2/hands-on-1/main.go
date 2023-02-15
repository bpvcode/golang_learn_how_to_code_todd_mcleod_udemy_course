package main

import "fmt"

const number int = 911

func main() {
	fmt.Println(number)         // Prints as int
	fmt.Printf("%T\n", number)  // Prints the type of number
	fmt.Printf("%b\n", number)  // Prints in binary
	fmt.Printf("%#X\n", number) // Prints in hexadecimal
}
