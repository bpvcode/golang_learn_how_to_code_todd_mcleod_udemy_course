package main

import "fmt"

func main() {
	fmt.Println("############################################# SECTION - CREATE AND ACCESS SLICE #############################################")

	// SLICE of type int
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}

	// Will print the SLICE from position 2 until the end (Position 2 is equal to value 3)
	fmt.Println(x[2:])
	// include the first not include the last
	fmt.Println(x[1:3])

	fmt.Println("############################################# SECTION - APPEND TO SLICE #############################################")

	// append returns []TYPE (returns a slice of the type of the slice passed in the first argument, in this case type x, so type int)
	x = append(x, 7, 8, 9, 10)
	fmt.Println(x)

	// APPEND a SLICE to a SLICE
	y := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	x = append(x, y...) // ... variatic parameters
	fmt.Println(x)

	fmt.Println("############################################# SECTION - DELETE FROM SLICE #############################################")

	z := x[4:]
	fmt.Println(z)

	//DELETE from a slice (Remove numbers 10 and 11 from slice x)
	x = append(x[:9], x[11:]...)
	fmt.Println(x)

	fmt.Println("############################################# SECTION - MAKE #############################################")

	//MAKE create a slice with a specific memory size allocations
	// Slice of type string with initial length 0 and maximum capacity 12
	k := make([]string, 0, 12)
	fmt.Println(k)
	fmt.Println(len(k)) // length
	fmt.Println(cap(k)) // capacity

	k = append(k, "|Hello|", "|Bob|", "|Hey|", "|Bob|", "|Oi|", "|Bob|")
	fmt.Println(k)
	fmt.Println(len(k))
	fmt.Println(cap(k))

	k = append(k, "|Xau|", "|Bob|", "|Bye|", "|Bob|")
	fmt.Println(k)
	fmt.Println(len(k))
	fmt.Println(cap(k))

	k = append(k, "|You're doing...|", "|GREAT!|")
	fmt.Println(k)
	fmt.Println(len(k))
	fmt.Println(cap(k))

	k = append(k, "|Check CAPACITY now...|", "|IT'S THE DOUBLE OF THE INITIAL SET ;)|")
	fmt.Println(k)
	fmt.Println(len(k))
	fmt.Println(cap(k))

	fmt.Println("############################################# SECTION - MULTIDIMENSIONAL SLICE #############################################")

	// Multi dimensional slice
	men := []string{"Bruno", "29"}
	women := []string{"Hydra", "30"}
	fmt.Println(men)
	fmt.Println(women)

	persons := [][]string{men, women}
	fmt.Println(persons)

	fmt.Println("############################################# SECTION - MAP #############################################")

	//MAP - [key]value
	m := map[string]int{
		"Bruno": 29,
		"Hydra": 30,
	}
	fmt.Println(m)
	fmt.Println(m["Bruno"])
	fmt.Println(m["Hydra"])

	// "comma ok" idiom -> Check if key exists on MAP
	v, ok := m["NO ENTRY"]
	fmt.Println(v)  // value is 0
	fmt.Println(ok) // is false, so key doesn't exist on map

	// if 'ok' is true
	if v, ok := m["Bruno"]; ok {
		fmt.Println("KEY MATCHES AN ENTRY ON MAP")
		fmt.Println(v)
	}

	// ADD NEW ITEM TO MAP
	m["Test"] = 10
	fmt.Println(m)
	fmt.Println(len(m))

	// DELETE FROM A MAP
	delete(m, "Test")
	fmt.Println(m)
	fmt.Println(len(m))

	// wrongKey := "MAL" change comments to see if deletes or no
	correctKey := "Bruno"

	if v, ok := m[correctKey]; ok {
		fmt.Println("VALUE", v)
		delete(m, correctKey)
		fmt.Println(m)
		fmt.Println(len(m))
	}

	// RANGE over a map
	for k, v := range m {
		fmt.Printf("NAME: %s\t AGE: %d\n", k, v)
	}

	// EXTRA - RANGE over a slice
	t := []int{2, 5, 6, 9, 7, 8, 4, 5, 6, 12, 3, 5, 8, 64}
	for k, v := range t {
		fmt.Printf("KEY: %v\t VALUE: %v\n", k, v)
	}
}
