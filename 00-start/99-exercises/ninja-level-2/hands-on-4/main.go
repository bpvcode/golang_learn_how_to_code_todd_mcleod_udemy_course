package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("%d\n%b\n%#X\n", x, x, x)

	y := x << 1
	fmt.Printf("%d\n%b\n%#X\n", y, y, y)
}
