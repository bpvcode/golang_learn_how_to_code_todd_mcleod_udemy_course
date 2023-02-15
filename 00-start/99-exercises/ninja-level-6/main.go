package main

import (
	"fmt"
)

// Defer a function will automatic run the function deferred when the surrounding function reaches the end or a go routine panics
// EXAMPLE: We open a communication channel, then we need to guarantee that closes at the end, so after we open, we can defer a function to close `defer closeChannel()`. This function will run at the end of the surrounding function although is called right after `openChannel()`

// FUNCTIONS can be used through all the module
// METHODS should be assigned to a TYPE, and only can be invoked by values of that TYPE

// interface{} -> every value implements the type 'interface{}', so any value could be input here or no values

type shipmentBasicInfo struct {
	trackingNumber string
	invoiceNumber  string
}

type geodis struct {
	shipmentBasicInfo
	printingLabels string
}

type gls struct {
	shipmentBasicInfo
	glsDescription string
}

type shipment interface {
	record()
	cancel()
	update()
}

func (order gls) getDescription() {
	fmt.Println("This is the description for GLS order ", order)
}

func (order geodis) getPrintingLabels() {
	fmt.Println("This are the printing labels for GEODIS order ", order)
}

func (s geodis) record() {
	fmt.Println("RECORD: ", s)
}

func (s geodis) cancel() {
	fmt.Println("CANCEL: ", s)
}

func (s geodis) update() {
	fmt.Println("UPDATE: ", s)
}

func (s gls) record() {
	fmt.Println("RECORD: ", s)
}

func (s gls) cancel() {
	fmt.Println("CANCEL: ", s)
}

func (s gls) update() {
	fmt.Println("UPDATE: ", s)
}

func recordShipment(s shipment) {
	switch s.(type) {
	case gls:
		fmt.Println("RECORD A GLS ORDER")
		s.record()
		fmt.Println("CANCEL A GLS ORDER")
		s.cancel()
		fmt.Println("UPDATE A GLS ORDER")
		s.update()
		//Assertion:
		fmt.Println(s.(gls).glsDescription)
	case geodis:
		fmt.Println("RECORD A GEODIS ORDER")
		s.record()
		fmt.Println("CANCEL A GEODIS ORDER")
		s.cancel()
		fmt.Println("UPDATE A GEODIS ORDER")
		s.update()
		//Assertion:
		fmt.Println(s.(geodis).printingLabels)
	}
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func sumEven(sumFunc func(xi ...int) int, si ...int) int {
	var evenNumbers []int
	for _, v := range si {
		if v%2 == 0 {
			evenNumbers = append(evenNumbers, v)
		}
	}
	return sumFunc(evenNumbers...)
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func factorialRecursion(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialRecursion(n-1)
}

func factorialLoop(n int) int {
	final := 1
	for i := n; i > 0; i-- {
		final *= i
	}
	return final
}

func main() {

	sb1 := shipmentBasicInfo{
		trackingNumber: "00001",
		invoiceNumber:  "INV00001",
	}

	sb2 := shipmentBasicInfo{
		trackingNumber: "00002",
		invoiceNumber:  "INV00002",
	}

	geo := geodis{
		sb1,
		"PRINTING LABELS",
	}

	gl := gls{
		sb2,
		"GLS DESCRIPTION",
	}

	recordShipment(geo)
	fmt.Println("#########################")
	recordShipment(gl)
	fmt.Println("#########################")
	fmt.Printf("%T\n", geo)
	fmt.Printf("%T\n", gl)

	//Anonymous function
	func() {
		fmt.Println("This is an anonymous function. There is no identifier!")
	}()

	func(s string) {
		fmt.Println("This is another anonymous function with arguments", s)
	}("YEY!")

	f := func(s string) {
		fmt.Println("This is a func expression", s)
	}
	f("-> ASSIGN A FUNCTION TO A VARIABLE")

	// Callback - pass a function as an argument to another function
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	total := sum(xi...)
	fmt.Println("SUM ->", total)

	totalEven := sumEven(sum, xi...)
	fmt.Println("SUM EVEN ->", totalEven)

	// Closure
	a := incrementor()
	fmt.Printf("%T\n", a)
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println("##################")
	b := incrementor()
	fmt.Printf("%T\n", b)
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

	// Recursion - When a functions calls it self
	// NOTE: FAVOR LOOPS OVER RECURSIONS (more complex to understand than loops, much memory usage)
	rec := factorialRecursion(4) // 4 * 3 * 2 * 1 * 1
	fmt.Println("FACTORIAL (RECURSION) DE 4:", rec)
	// LOOP INSTEAD RECURSION
	loop := factorialLoop(4) // 4 * 3 * 2 * 1
	fmt.Println("FACTORIAL (LOOP) DE 4:", loop)
}
