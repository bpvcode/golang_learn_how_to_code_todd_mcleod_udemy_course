package main

import "fmt"

func main() {
	x := 2 == 1
	y := 2 >= 1
	z := 1 <= 1
	k := "true" != "false"
	l := 2 > 3
	m := 3 < 3

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
}
