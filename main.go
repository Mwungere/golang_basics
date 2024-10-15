package main

import "fmt"

var someName = "Mario"


func main(){

	//variables

	//strings
	var nameOne string = "Ninja"
	var nameTwo = "Luigi"
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree);

	nameOne = "peach"
	nameThree = "bowser"
	fmt.Println(nameOne, nameTwo, nameThree);

	nameFour :="yoshi"

	fmt.Println(nameOne, nameTwo, nameThree, nameFour, someName);

	//ints
	var ageOne int = 15
	var ageTwo = 16
	var ageThree int
	fmt.Println(ageOne, ageTwo);

	ageThree = 18
	ageFour:=20
	fmt.Println(ageOne, ageTwo, ageThree, ageFour);

	//bits and memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint8 = 255
	fmt.Println(numOne, numTwo, numThree);

	//floats
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 256478920348208384924.3453
	scoreThree := 923423.4
	fmt.Println(scoreOne, scoreTwo, scoreThree)

}
