package main

import "fmt"

func main() {
	// HANDS-ON 1
	fmt.Println("######### HANDS-ON 1 #########")
	// ARRAY
	fmt.Println("*** ARRAY")
	array := [5]int{1, 2, 3, 4, 5}

	for _, v := range array {
		fmt.Println("VALUE: ", v)
	}

	fmt.Printf("TYPE: %T\n", array)

	// HANDS-ON 2
	fmt.Println("######### HANDS-ON 2 #########")
	// SLICE
	fmt.Println("*** SLICE")
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range slice {
		fmt.Println("VALUE: ", v)
	}

	fmt.Printf("TYPE: %T\n", slice)

	// SLICE - with make
	fmt.Println("*** SLICE WITH MAKE")

	sliceMake := make([]int, 0, 5)
	fmt.Println(sliceMake)
	fmt.Println(len(sliceMake))
	fmt.Println(cap(sliceMake))

	sliceMake = append(sliceMake, 1, 2, 3, 4, 5)
	fmt.Println(sliceMake)
	fmt.Println(len(sliceMake))
	fmt.Println(cap(sliceMake))

	sliceMake = append(sliceMake, 6)
	fmt.Println(sliceMake)
	fmt.Println(len(sliceMake))
	fmt.Println(cap(sliceMake)) // Doubles the initial set

	fmt.Printf("TYPE: %T\n", sliceMake)

	// HANDS-ON 3
	fmt.Println("######### HANDS-ON 3 #########")
	// SLICING
	newSlice := slice[3:]      // 4-10
	newSlice1 := slice[4:]     // 5-10
	newSlice2 := slice[:8]     // 1-8
	newSlice3 := slice[1:7]    // 2-7
	newSlice4 := sliceMake[:3] // 1-3

	fmt.Println(newSlice)
	fmt.Println(newSlice1)
	fmt.Println(newSlice2)
	fmt.Println(newSlice3)
	fmt.Println(newSlice4)

	// HANDS-ON 4
	fmt.Println("######### HANDS-ON 4 #########")
	slice = append(slice, 11)
	fmt.Println(slice)
	fmt.Printf("TYPE: %T\n", slice)

	slice = append(slice, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	fmt.Println(slice)

	sliceToAppend := []int{21, 22, 23, 24, 25}
	slice = append(slice, sliceToAppend...)
	fmt.Println(slice)
	fmt.Printf("TYPE: %T\n", slice)

	// HANDS-ON 5
	fmt.Println("######### HANDS-ON 5 #########")
	slice = append(slice[:9], slice[19:]...)
	fmt.Println(slice)

	// HANDS-ON 6
	fmt.Println("######### HANDS-ON 6 #########")
	make := make([]string, 0, 50)
	fmt.Println(make)
	fmt.Println(len(make))
	fmt.Println(cap(make))

	states := map[string]string{
		"AL": "Alabama",
		"AK": "Alaska",
		"AZ": "Arizona",
		"AR": "Arkansas",
		"CA": "California",
		"CO": "Colorado",
		"CT": "Connecticut",
		"DE": "Delaware",
		"FL": "Florida",
		"GA": "Georgia",
		"HI": "Hawaii",
		"ID": "Idaho",
		"IL": "Illinois",
		"IN": "Indiana",
		"IA": "Iowa",
		"KS": "Kansas",
		"KY": "Kentucky",
		"LA": "Louisiana",
		"ME": "Maine",
		"MD": "Maryland",
		"MA": "Massachusetts",
		"MI": "Michigan",
		"MN": "Minnesota",
		"MS": "Mississippi",
		"MO": "Missouri",
		"MT": "Montana",
		"NE": "Nebraska",
		"NV": "Nevada",
		"NH": "New Hampshire",
		"NJ": "New Jersey",
		"NM": "New Mexico",
		"NY": "New York",
		"NC": "North Carolina",
		"ND": "North Dakota",
		"OH": "Ohio",
		"OK": "Oklahoma",
		"OR": "Oregon",
		"PA": "Pennsylvania",
		"RI": "Rhode Island",
		"SC": "South Carolina",
		"SD": "South Dakota",
		"TN": "Tennessee",
		"TX": "Texas",
		"UT": "Utah",
		"VT": "Vermont",
		"VA": "Virginia",
		"WA": "Washington",
		"WV": "West Virginia",
		"WI": "Wisconsin",
		"WY": "Wyoming",
	}

	for _, v := range states {
		make = append(make, v)
	}
	fmt.Println(make)
	fmt.Println(len(make))
	fmt.Println(cap(make))

	// HANDS-ON 7
	fmt.Println("######### HANDS-ON 7 #########")
	sliceOfSlice := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo James"}}
	fmt.Println(sliceOfSlice)

	for i, v := range sliceOfSlice {
		fmt.Printf("INDEX: %v \t VALUE: %v\n", i, v)
		for index, value := range v {
			fmt.Printf("\t INDEX: %v \t VALUE: %v\n", index, value)
		}
	}

	// HANDS-ON 8
	fmt.Println("######### HANDS-ON 8 #########")
	newMap := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	fmt.Println(newMap)

	for k, v := range newMap {
		fmt.Printf("KEY: %v \t VALUE: %v\n", k, v)
		for index, value := range v {
			fmt.Printf("\t INDEX: %v \t VALUE: %v\n", index, value)
		}
	}

	// HANDS-ON 9
	fmt.Println("######### HANDS-ON 9 #########")
	newMap["bruno_vilar"] = []string{"computer science", "golang", "javascript", "devops"}

	for k, v := range newMap {
		fmt.Printf("KEY: %v\n", k)
		fmt.Println("VALUES:")
		for _, v2 := range v {
			fmt.Printf("\t%v\n", v2)
		}
	}

	// HANDS-ON 10
	fmt.Println("######### HANDS-ON 10 #########")
	delete(newMap, "no_dr")

	for k, v := range newMap {
		fmt.Printf("KEY: %v\n", k)
		fmt.Println("VALUES:")
		for _, v2 := range v {
			fmt.Printf("\t%v\n", v2)
		}
	}
}
