package main

import "fmt"


func main(){

	//arrays

	var ages [3]int = [3]int{20,23,26}
	var ages2 = [3]int{20,30,40}
	names:= [4]string{"Bob","Mark","John","Sally"}
	names[1] = "Luigi"
	fmt.Println(ages, len(ages2), names)

	//slices

	var scores[]int = []int{10,20,30,40,50,60,70,80,90,100,200}
	scores = append(scores, 300)
	fmt.Println(scores, len(scores))

	//slice ranges

	rangeOne := scores[0:5]
	fmt.Println(rangeOne)
	rangeTwo := scores[5:10]
	fmt.Println(rangeTwo)
	rangeThree := scores[5:]
	fmt.Println(rangeThree)
	rangeFour := scores[:5]
	fmt.Println(rangeFour)
}
