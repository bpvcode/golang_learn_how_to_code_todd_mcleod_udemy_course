package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	foo()
	count(20)
}

func foo() {
	fmt.Println("I'm in foo")
}

func count(x int) {
	var counter = 0
	for i := 0; i < x; i++ {
		counter++
		fmt.Println("Counter ", counter)
	}
}
