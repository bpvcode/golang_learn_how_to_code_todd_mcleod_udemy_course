package main

import "fmt"

func main() {
	a := 99
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b := &a               // Points to the memory location of "a"
	fmt.Println(b)        // memory ADDRESS location
	fmt.Printf("%T\n", b) // Type -> Pointer to an int

	fmt.Println(*b)  // The value stored in memory address location - Dereference the address
	fmt.Println(*&a) // Dereference the value stored in the memory location address of "a"

	*b = 100       // Change the value stored in the address
	fmt.Println(a) // initial value of a changed

	c := 1000
	fmt.Println(&c)
	fmt.Println(c)

	c++
	fmt.Println(&c)
	fmt.Println(c)

	// MUTATE - change the value stored in an address

}
