package main

import (
	"fmt"
)

type person struct {
	firstName        string
	lastName         string
	favoriteIceCream []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	// HANDS-ON 1
	fmt.Println("######### HANDS-ON 1 #########")
	p1 := person{
		firstName:        "Bruno",
		lastName:         "Vilar",
		favoriteIceCream: []string{"Manga", "Chocolate"},
	}

	p2 := person{
		firstName:        "Hydra",
		lastName:         "Smith",
		favoriteIceCream: []string{"Vanilla", "Peach"},
	}

	var persons []person = []person{p1, p2}

	for _, v := range persons {
		fmt.Println("PERSON: ", v.firstName+" "+v.lastName)
		fmt.Println("\tFavorite ice cream flavours:")
		for _, value := range v.favoriteIceCream {
			fmt.Printf("\t%v\n", value)
		}
	}

	// HANDS-ON 2
	fmt.Println("######### HANDS-ON 2 #########")
	newMap := map[string]person{}
	newMap[p1.lastName] = p1
	newMap[p2.lastName] = p2

	fmt.Println(newMap)

	for k, v := range newMap {
		fmt.Println("PERSON: ", v.firstName+" "+k)
		fmt.Println("\tFavorite ice cream flavours:")
		for _, value := range v.favoriteIceCream {
			fmt.Printf("\t%v\n", value)
		}
	}

	// HANDS-ON 3
	fmt.Println("######### HANDS-ON 3 #########")
	vehicle1 := truck{
		vehicle: vehicle{
			doors: 5,
			color: "Black",
		},
		fourWheel: true,
	}

	vehicle2 := sedan{
		vehicle: vehicle{
			doors: 3,
			color: "Red",
		},
		luxury: true,
	}

	fmt.Println("TRUCK: ", vehicle1)
	fmt.Println("SEDAN: ", vehicle2)

	fmt.Println("TRUCK SINGLE VALUE: ", vehicle1.vehicle)
	fmt.Println("SEDAN SINGLE VALUE: ", vehicle2.doors)

	// HANDS-ON 4
	fmt.Println("######### HANDS-ON 4 #########")
	vehicle3 := struct {
		vehicle
		extras_price    map[string]int
		characteristics []string
	}{
		vehicle: vehicle{
			doors: 5,
			color: "Black",
		},
		extras_price: map[string]int{
			"auto_pilot":    2000,
			"leader":        500,
			"person_sensor": 1000,
		},
		characteristics: []string{
			"Luxury",
			"Sport",
			"City",
		},
	}

	fmt.Println(vehicle3)

	fmt.Println("EXTRAS PRICES: ")
	for k, v := range vehicle3.extras_price {
		fmt.Printf("\t%v : %v€\n", k, v)
	}

	fmt.Println("CHARACTERISTICS: ")
	for _, v := range vehicle3.characteristics {
		fmt.Printf("\t%v\n", v)
	}
}
