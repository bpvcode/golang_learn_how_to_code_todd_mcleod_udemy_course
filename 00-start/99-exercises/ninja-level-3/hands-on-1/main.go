package main

import "fmt"

func main() {
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}

	fmt.Println("#############################")

	y := 1
	for y <= 10 {
		fmt.Println(y)
		y++
	}
}
