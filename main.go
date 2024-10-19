package main

import (
	"fmt"
	"math"
)



func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func sayGreeting (n string) {
	fmt.Println("Hello " + n)
}

func cycleNames (n []string, f func(string)){
	for _, name := range n{
		f(name)
	}
}

func sayBye (n string) {
	fmt.Println("Bye " + n)
}

func main() {

	names := []string {"Job", "John", "Jake", "Duke"}

	//funcitons
	sayGreeting("John")
	sayBye("John")
	cycleNames(names, sayGreeting);
	cycleNames(names, sayBye);
	fmt.Printf("Area = %0.2f",circleArea(5.3))

	
}
