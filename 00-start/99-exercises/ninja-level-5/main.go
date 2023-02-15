package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Bruno",
		age:   29,
	}
	fmt.Println(p1)
}
