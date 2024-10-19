package main

import "fmt"

// maps are like dictionaries in python or objects
//in a map all keys must be of the same datatype and all values must be of the same datatype

var score = 99.5
func main() {
	menu := map[string]float64{
		"soup": 4.99,
		"pie":7.99,
		"salad": 6.99,
		"toffee padding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	for k, v := range menu{
		fmt.Println(k,"-",v)
	}

	// ints as key type
	phoneBook := map[int]string{
		1234567: "Mario",
		5435353: "Luigi",
		3945854: "Peach",
	}

	fmt.Println(phoneBook)
	fmt.Println(phoneBook[5435353])

	phoneBook[5435353] = "Bowser"
	for k, v := range phoneBook{
		fmt.Println(k,"-",v)
	}
}
