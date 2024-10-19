package main

import (
	"fmt"
	"strings"
)


//multiple return values

func getInitials (n string) (string, string){
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func main() {
fn1, sn1 := getInitials("Tifa Lockhart")
fn2, sn2 := getInitials("John Doe")
fn3, sn3 := getInitials("barret")

fmt.Println(fn1,sn1)
fmt.Println(fn2,sn2)
fmt.Println(fn3, sn3)


	
}
