package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

type person struct {
	name string
}

type human interface {
	speak() string
}

func (p *person) speak() string {
	return fmt.Sprintf("Hello, is %s speaking...\n", p.name)
}

func saySomething(h human) {
	fmt.Println(h.speak())
}

func main() {
	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("GO ROUTINES: ", runtime.NumGoroutine())

	// HANDS-ON 1
	fmt.Println("######### HANDS-ON 1 #########")
	counter := 0
	wg.Add(2)
	go func() {
		counter++
		wg.Done()
	}()
	go func() {
		counter++
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(counter)

	// HANDS-ON 2
	fmt.Println("######### HANDS-ON 2 #########")
	p1 := person{name: "Bruno"}
	saySomething(&p1)

	// HANDS-ON 3 && 4
	fmt.Println("######### HANDS-ON 3 && 4 #########")
	var wg1 sync.WaitGroup
	var mu sync.Mutex
	wg1.Add(100)
	total := 0
	for i := 1; i <= 100; i++ {
		go func() {
			mu.Lock()
			v := total
			v++
			total = v
			mu.Unlock()
			wg1.Done()
		}()
	}
	wg1.Wait()
	fmt.Println("TOTAL: ", total)

	// HANDS-ON 5
	fmt.Println("######### HANDS-ON 5 #########")
	var wg2 sync.WaitGroup
	wg2.Add(100)
	var total1 int64
	for i := 1; i <= 100; i++ {
		go func() {
			atomic.AddInt64(&total1, 1)
			runtime.Gosched()
			wg2.Done()
		}()
	}
	wg2.Wait()
	fmt.Println("TOTAL: ", atomic.LoadInt64(&total1))

	// HANDS-ON 6
	fmt.Println("######### HANDS-ON 6 #########")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
