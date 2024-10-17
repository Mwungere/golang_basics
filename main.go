package main

import (
	"fmt"
	"sort"
	"strings"
)


func main(){

	greetings := "Hello there friends!"
	fmt.Println(strings.Contains(greetings, "friends"))
	fmt.Println(strings.Contains(greetings, "friends"))
	fmt.Println(strings.ToUpper(greetings))
	fmt.Println(strings.Index(greetings, "ll"))
	fmt.Println(strings.Split(greetings, " "))

	// for a slice of integers
	ages := []int{100, 20, 130, 40, 15, 60, 17, 80, 9, 0}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 400)
	fmt.Println("Found 40 at index ", index)

	// for a slice of strings
	names := []string{"Charlie", "Alice", "Luigi", "Dave","Bob"}
	sort.Strings(names)

	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "Bob"))
}
