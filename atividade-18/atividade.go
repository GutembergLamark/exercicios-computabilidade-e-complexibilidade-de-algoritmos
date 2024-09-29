package main

import "fmt"

type AFD struct {
	state string
}

func (afd *AFD) transition(char rune) {
	switch afd.state {
	case "q00":
		if char == '0' {
			afd.state = "q10"
		} else if char == '1' {
			afd.state = "q01"
		}
	case "q01":
		if char == '0' {
			afd.state = "q11"
		} else if char == '1' {
			afd.state = "q00"
		}
	case "q10":
		if char == '0' {
			afd.state = "q00"
		} else if char == '1' {
			afd.state = "q11"
		}
	case "q11":
		if char == '0' {
			afd.state = "q01"
		} else if char == '1' {
			afd.state = "q10"
		}
	}
}

func (afd *AFD) isAccepted(input string) bool {
	afd.state = "q00" 
	for _, char := range input {
		afd.transition(char)
	}
	return afd.state == "q11" 
}

func main() {
	afd := AFD{}
	fmt.Println(afd.isAccepted("01"))   
	fmt.Println(afd.isAccepted("1100")) 
	fmt.Println(afd.isAccepted("10"))   
	fmt.Println(afd.isAccepted("0000")) 
	fmt.Println(afd.isAccepted("1111")) 
}
