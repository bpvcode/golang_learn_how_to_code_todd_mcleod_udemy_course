package main

import "fmt"

func main() {
	var alphabet = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	i := 0
	for i < 26 {
		j := 0
		fmt.Printf("############## %v\n", i)
		for j < 3 {
			fmt.Printf("%v\n", alphabet[i])
			j++
		}
		i++
	}
}
