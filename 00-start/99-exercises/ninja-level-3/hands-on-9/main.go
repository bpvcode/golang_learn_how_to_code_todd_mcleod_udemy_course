package main

import "fmt"

func main() {

	const favSport string = "Surf"

	switch favSport {
	case "Football":
		fmt.Println("FOOTBALL")
	case "Tennis":
		fmt.Println("TENNIS")
	case "Gym":
		fmt.Println("GYM")
	case "Surf":
		fmt.Println("SURF")
	default:
		fmt.Println("YOU DON'T LIKE SPORTS")
	}

}
