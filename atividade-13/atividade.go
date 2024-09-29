package main

import "fmt"

type AFD struct {
	state string
}

func NewAFD() *AFD {
	return &AFD{state: "q0"}
}

func (afd *AFD) transition(input rune) {
	switch afd.state {
	case "q0":
		if input == '1' {
			afd.state = "q1"
		} else if input == '0' {
			afd.state = "q-1"
		}
	case "q1":
		if input == '1' {
			afd.state = "q2"
		} else if input == '0' {
			afd.state = "q0"
		}
	case "q-1":
		if input == '1' {
			afd.state = "q0"
		} else if input == '0' {
			afd.state = "q-2"
		}
	case "q2":
		if input == '1' {
			afd.state = "q3"
		} else if input == '0' {
			afd.state = "q1"
		}
	case "q-2":
		if input == '1' {
			afd.state = "q-1"
		} else if input == '0' {
			afd.state = "q-3"
		}
	}
}

func (afd *AFD) isAccepted(input string) bool {
	afd.state = "q0"
	for _, c := range input {
		afd.transition(c)
	}


	return afd.state == "q1" || afd.state == "q2" || afd.state == "q3"
}

func main() {
	afd := NewAFD()
	fmt.Println(afd.isAccepted("110"))  
	fmt.Println(afd.isAccepted("101"))  
	fmt.Println(afd.isAccepted("11100"))
	fmt.Println(afd.isAccepted("100"))  
	fmt.Println(afd.isAccepted("111"))  
}
