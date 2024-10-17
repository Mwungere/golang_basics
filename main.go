package main

import (
	"fmt"
	
)


func main(){

	//loops

	// similar to while loop
	x := 0

	for x < 5 {
		fmt.Println("The value of x is: ", x)
		x++
	}

	// for loop

	for y := 0; y<5; y++ {
		fmt.Println("The value of y is: ", y)
	}

	names := []string{"John", "Paul", "George", "Ringo"}

	for index, value := range names {
		fmt.Printf("The value at index %v is %v\n", index, value)
	}

	//if you do not need to use the index or value just replace any with an underscore

	for _, value := range names {
		fmt.Printf("The value is %v\n", value)
	}
}
