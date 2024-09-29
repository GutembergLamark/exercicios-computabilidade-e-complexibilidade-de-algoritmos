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
	switch afd.state {
	case "q0":
		if c == '0' {
			afd.state = "q1"
		} else if c == '1' {
			afd.state = "q0"
		}
	case "q1":
		if c == '0' {
			afd.state = "q0"
		} else if c == '1' {
			afd.state = "q1"
		}
	}
}

func (afd *AFD) isAccepted(input string) bool {
	afd.state = "q0" 
	for _, c := range input {
		afd.transition(c)
	}
	return afd.state == "q0"
}

func main() {
	afd := NewAFD()
	fmt.Println(afd.isAccepted("110"))  
	fmt.Println(afd.isAccepted("101"))  
	fmt.Println(afd.isAccepted("1001")) 
	fmt.Println(afd.isAccepted("0"))    
	fmt.Println(afd.isAccepted("00"))   
}
