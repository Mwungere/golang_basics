package main

import "fmt"


func main(){
	age:= 20;
	name:= "John";

	//Print
	fmt.Print("Hello, ")
	fmt.Print("world! \n")
	fmt.Print("New line \n")

	//Println
	fmt.Println("Hello Ninjas!")
	fmt.Println("Goodbye Ninjas!")
	fmt.Println("My age is", age, "and my name is", name)

	//Printf (formatted strings) %_ = format specifier 
	fmt.Printf("My age is %v and my name is %v \n", age, name)
	fmt.Printf("My age is %v and my name is %q \n", age, name)
	fmt.Printf("Age is of type %T \n", age)
	fmt.Printf("You scored %f points! \n", 225.55)
	fmt.Printf("You scored %0.1f points! \n", 225.55)


}
