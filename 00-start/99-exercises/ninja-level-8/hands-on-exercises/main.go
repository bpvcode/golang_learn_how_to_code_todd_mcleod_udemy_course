package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type User struct {
	UserName string   `json:"username"`
	Password string   `json:"password"`
	Age      uint8    `json:"age"`
	Sports   []string `json:"sports"`
}

var users []User = []User{
	{
		UserName: "ze",
		Password: "123456",
		Age:      20,
		Sports:   []string{"surf"},
	},
	{
		UserName: "hugo",
		Password: "password",
		Age:      15,
		Sports:   []string{"ping-pong", "tennis"},
	},
	{
		UserName: "andreia",
		Password: "changeMe",
		Age:      10,
		Sports:   []string{"surf", "gym", "yoga"},
	},
	{
		UserName: "bea",
		Password: "999999",
		Age:      76,
		Sports:   []string{"golf", "skate"},
	},
}

var users01 []User

var usersMockData []byte = []byte(`[{"username": "user0", "password": "123456", "age": 20}, {"username": "user1", "password": "password", "age": 21}, {"username": "user2", "password": "changeMe", "age": 22}, {"username": "user3", "password": "999999", "age": 23}, {"username": "user4", "password": "777777", "age": 24}]`)

func main() {
	// HANDS-ON 1
	fmt.Println("######### HANDS-ON 1 #########")
	jsn, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsn))

	// HANDS-ON 2
	fmt.Println("######### HANDS-ON 2 #########")
	fmt.Println(users01)
	err = json.Unmarshal(usersMockData, &users01)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users01)

	for i, v := range users01 {
		fmt.Println("USER ", i)
		fmt.Printf("\tUsername: %v\n", v.UserName)
		fmt.Printf("\tPassword: %v\n", v.Password)
		fmt.Printf("\tAge: %v\n", v.Age)
	}

	// HANDS-ON 3
	fmt.Println("######### HANDS-ON 3 #########")
	err = json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(err)
	}

	// HANDS-ON 4 && 5
	fmt.Println("######### HANDS-ON 4 && 5 #########")
	sort.Slice(users[:], func(i, j int) bool {
		return users[i].Age < users[j].Age
	})
	fmt.Println("USERS SORTED BY AGE")
	fmt.Println(users)

	sort.Slice(users[:], func(i, j int) bool {
		return users[i].UserName < users[j].UserName
	})
	fmt.Println("USERS SORTED BY USERNAME")
	fmt.Println(users)

	for _, v := range users {
		sort.Strings(v.Sports)
	}
	fmt.Println("USERS SORTED BY SPORTS")
	fmt.Println(users)
}
