package main

import (
	"fmt"
)

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func gen1(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func receive1(cd, q <-chan int) {
	for {
		select {
		case v := <-cd:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}

func main() {
	// HANDS-ON 1
	fmt.Println("######### HANDS-ON 1 #########")
	c := make(chan int, 1) // buffered channel
	c <- 42
	fmt.Println(<-c)

	d := make(chan int) //using function literal
	go func() {
		d <- 43
	}()
	fmt.Println(<-d)

	// HANDS-ON 2
	fmt.Println("######### HANDS-ON 2 #########")
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	// HANDS-ON 3
	fmt.Println("######### HANDS-ON 3 #########")
	cs2 := gen()

	receive(cs2)
	fmt.Println("about to exit")

	// HANDS-ON 4
	fmt.Println("######### HANDS-ON 4 #########")
	q := make(chan int)
	cd := gen1(q)

	receive1(cd, q)
	fmt.Println("about to exit")

	// HANDS-ON 5
	fmt.Println("######### HANDS-ON 5 #########")
	cc := make(chan int)

	go func() {
		cc <- 42
		cc <- 23
	}()

	v, ok := <-cc
	fmt.Println(v, ok)

	v, ok = <-cc
	fmt.Println(v, ok)
	close(cc)

	v, ok = <-cc
	fmt.Println(v, ok)

	// HANDS-ON 6
	fmt.Println("######### HANDS-ON 6 #########")
	nc := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			nc <- i
		}
		close(nc)
	}()

	for v := range nc {
		fmt.Println(v)
	}
	fmt.Println("about to exit")

	// HANDS-ON 7
	fmt.Println("######### HANDS-ON 7 #########")
	nc1 := make(chan int)

	for i := 1; i <= 10; i++ {
		go func() {
			for i := 1; i <= 10; i++ {
				nc1 <- i
			}
		}()
	}

	for i := 1; i <= 100; i++ {
		fmt.Println(i, <-nc1)
	}
	close(nc1)
	fmt.Println("about to exit")
}
