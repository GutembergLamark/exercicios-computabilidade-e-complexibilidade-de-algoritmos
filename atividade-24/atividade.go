package main

import "fmt"

type AFD struct {
	state  string
	rotate bool
}

func NewAFD() *AFD {
	return &AFD{
		state:  "q0",
		rotate: false,
	}
}

func (afd *AFD) transition(c rune) {
	switch afd.state {
	case "q0":
		if c == '0' {
			afd.state = "q1"
		}
	case "q1":
		if c == '1' {
			afd.state = "q2"
		} else {
			afd.state = "q0"
		}
	case "q2":
		if c == '0' && !afd.rotate {
			afd.state = "q0"
			afd.rotate = true
		} else if c == '0' && afd.rotate {
			afd.state = "q3" 
		} else if c == '1' {
			afd.state = "q0"
		}
	}
}

func (afd *AFD) isAccepted(input string) bool {
	afd.state = "q0"
	afd.rotate = false
	for _, c := range input {
		afd.transition(c)
	}
	return afd.state == "q3"
}

func main() {
	afd := NewAFD()
	fmt.Println(afd.isAccepted("010010"))  
	fmt.Println(afd.isAccepted("010"))     
	fmt.Println(afd.isAccepted("0101010")) 
	fmt.Println(afd.isAccepted("101010"))  
	fmt.Println(afd.isAccepted("111"))     
	fmt.Println(afd.isAccepted("000"))     
}
