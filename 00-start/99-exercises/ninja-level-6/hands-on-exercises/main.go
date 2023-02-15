package main

import (
	"fmt"
	"math"
)

func foo() int {
	return 1
}

func bar() (int, string) {
	return 100, "€"
}

func sumVariadicParam(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func sumSlice(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func showDefer(n int) {
	defer func() {
		fmt.Println("DEFER FUNCTION - PRINTS AT THE END OF THE FUNCTION ALWAYS")
	}()
	if n < 0 {
		fmt.Println("PRINTS FIRST")
	}
}

type animal struct {
	firstName string
	lastName  string
	age       uint
}

func (a animal) speak() {
	fmt.Println("Hello, I'm ", a.firstName, a.lastName, "and i'm an animal with", a.age, "years old!")
}

type square struct {
	side float64
}

type circle struct {
	diameter float64
}

type shape interface {
	calcArea() float64
}

func (s square) calcArea() float64 {
	return s.side * s.side
}

func (c circle) calcArea() float64 {
	raio := c.diameter / 2
	return math.Pi * (raio * raio)
}

func printArea(s1 shape) {
	fmt.Println(s1.calcArea())
}

func returnFunc() func() int {
	return func() int {
		return 1
	}
}

func sumAllNumbersGreaterThan(x []int, s int) int {
	total := 0
	for _, v := range x {
		if v > s {
			total += v
		}
	}

	return total
}

func sumAllEvenNumbersGreaterThan(f func(x []int, s int) int, xo []int, so int) int {
	var total []int
	for _, v := range xo {
		if v%2 == 0 {
			total = append(total, v)
		}
	}

	return f(total, so)
}

func inc() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}

func main() {

	// HANDS-ON 1
	fmt.Println("######### HANDS-ON 1 #########")
	fmt.Println(foo())
	fmt.Println(bar())

	// HANDS-ON 2
	fmt.Println("######### HANDS-ON 2 #########")
	sum := sumVariadicParam([]int{100, 1000}...)
	fmt.Println(sum)

	sum1 := sumSlice([]int{100, 1000})
	fmt.Println(sum1)

	// HANDS-ON 3
	fmt.Println("######### HANDS-ON 3 #########")
	showDefer(-1)
	showDefer(1)

	// HANDS-ON 4
	fmt.Println("######### HANDS-ON 4 #########")
	a := animal{
		firstName: "Pepper",
		lastName:  "Cookie",
		age:       3,
	}
	a.speak()

	// HANDS-ON 5
	fmt.Println("######### HANDS-ON 5 #########")
	circ := circle{
		diameter: 10,
	}

	sqr := square{
		side: 5,
	}

	printArea(circ)
	printArea(sqr)

	// HANDS-ON 6
	fmt.Println("######### HANDS-ON 6 #########")
	func(n int) {
		for ; n <= 10; n++ {
			fmt.Println(n)
		}
	}(5)

	// HANDS-ON 7
	fmt.Println("######### HANDS-ON 7 #########")
	h := func() {
		fmt.Println("TEST")
	}
	h()

	// HANDS-ON 8
	fmt.Println("######### HANDS-ON 8 #########")
	t := returnFunc()
	fmt.Println(t())

	// HANDS-ON 9
	fmt.Println("######### HANDS-ON 9 #########")
	sli := []int{10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 100}
	fmt.Println(sumAllEvenNumbersGreaterThan(sumAllNumbersGreaterThan, sli, 50))
	fmt.Println(sumAllEvenNumbersGreaterThan(sumAllNumbersGreaterThan, sli, 80))
	fmt.Println(sumAllEvenNumbersGreaterThan(sumAllNumbersGreaterThan, sli, 90))
	fmt.Println(sumAllEvenNumbersGreaterThan(sumAllNumbersGreaterThan, sli, 5))

	// HANDS-ON 10
	fmt.Println("######### HANDS-ON 10 #########")
	kla := inc()
	klb := inc()

	fmt.Println(kla())
	fmt.Println(kla())
	fmt.Println(kla())
	fmt.Println(kla())
	fmt.Println(klb())
	fmt.Println(klb())
	fmt.Println(klb())
	fmt.Println(klb())

}
