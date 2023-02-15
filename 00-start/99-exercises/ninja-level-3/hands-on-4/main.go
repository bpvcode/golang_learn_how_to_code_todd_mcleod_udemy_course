package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1994
	year, _, _ := time.Now().Date()
	for {
		fmt.Println(i)
		i++
		if i > year {
			break
		}
	}
}
