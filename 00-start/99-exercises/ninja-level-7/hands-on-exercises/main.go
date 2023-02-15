package main

import "fmt"

type person struct {
	firstName string
}

func (p *person) changedMe(name string) {
	p.firstName = name
}

func main() {

	// HANDS-ON 1
	fmt.Println("######### HANDS-ON 1 #########")
	a := 50
	fmt.Println(&a) // Print the address

	// HANDS-ON 2
	fmt.Println("######### HANDS-ON 2 #########")
	p1 := person{
		firstName: "Bruno",
	}
	fmt.Println(p1)

	p1.changedMe("NEW NAME")
	fmt.Println(p1)

}
