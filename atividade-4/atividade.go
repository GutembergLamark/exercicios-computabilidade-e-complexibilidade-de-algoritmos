package main

import (
	"fmt"
)

type AFD struct {
	state string
}

func (a *AFD) transition(char rune) {
	if a.state == "q0" {
		if char == '0' {
			a.state = "q1"
		}
	}
}

func (a *AFD) isAccepted(input string) bool {
	a.state = "q0" 
	for _, char := range input {
		a.transition(char)
	}
	return a.state == "q1" 
}

func main() {
	afd := AFD{}
	fmt.Println(afd.isAccepted("1"))    
	fmt.Println(afd.isAccepted("0"))    
	fmt.Println(afd.isAccepted("111"))  
	fmt.Println(afd.isAccepted("101"))  
	fmt.Println(afd.isAccepted("1100")) 
}
