package main

import "fmt"

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
		}
	case "q1":
		if c == 'b' {
			afd.state = "q2"
		} else if c == 'a' {
			afd.state = "q1"
		}
	case "q2":
		if c == 'a' {
			afd.state = "q3"
		}
	case "q3":
		
	}
}

func (afd *AFD) isAccepted(input string) bool {
	afd.state = "q0"
	for _, c := range input {
		afd.transition(c)
	}
	return afd.state == "q2"
}

func main() {
	afd := NewAFD()
	fmt.Println(afd.isAccepted("ab"))     
	fmt.Println(afd.isAccepted("aab"))    
	fmt.Println(afd.isAccepted("abab"))   
	fmt.Println(afd.isAccepted("baab"))   
	fmt.Println(afd.isAccepted("bbaab"))  
	fmt.Println(afd.isAccepted("baabab")) 
}
