package main

import (
	"fmt"
)

type AFD struct {
	state string
}

func NewAFD() *AFD {
	return &AFD{state: "q0"} 
}

func (afd *AFD) transition(c rune) {
	if afd.state == "q0" {
		if c == 'a' {
			afd.state = "q1" 
		}
	} else if afd.state == "q1" {
		if c == 'a' {
			afd.state = "q0" 
		}
	}
}

func (afd *AFD) isAccepted(input string) bool {
	afd.state = "q0" 
	for _, c := range input {
		afd.transition(c)
	}
	return afd.state == "q1" 
}

func main() {
	afd := NewAFD()
	fmt.Println(afd.isAccepted("a"))   
	fmt.Println(afd.isAccepted("aa"))  
	fmt.Println(afd.isAccepted("ab"))  
	fmt.Println(afd.isAccepted("bba")) 
	fmt.Println(afd.isAccepted("bb"))  
}
