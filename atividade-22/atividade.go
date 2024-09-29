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
		if c == 'a' {
			afd.state = "q1"
		} else if c == 'b' {
			afd.state = "q2"
		}
	case "q1":
		if c == 'a' {
			afd.state = "q2"
		} else if c == 'b' {
			afd.state = "q0"
		}
	case "q2":
		if c == 'a' {
			afd.state = "q0"
		} else if c == 'b' {
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
	fmt.Println(afd.isAccepted("aaabbb") 
	fmt.Println(afd.isAccepted("aab"))    
	fmt.Println(afd.isAccepted("aabb"))   
	fmt.Println(afd.isAccepted("aaa"))    
	fmt.Println(afd.isAccepted("bbb"))    
	fmt.Println(afd.isAccepted("ababab")) 
}
