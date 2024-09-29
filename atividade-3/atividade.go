package main

import (
	"fmt"
)

type AFD struct {
	state string
}

func (a *AFD) transition(char rune) {
	switch a.state {
	case "q0":
		if char == '1' {
			a.state = "q1"
		}
	case "q1":
		if char == '1' {
			a.state = "q2"
		}
	case "q2":
		if char == '1' {
			a.state = "q3"
		}
	}
}

func (a *AFD) isAccepted(input string) bool {
	a.state = "q0"
	for _, char := range input {
		a.transition(char)
	}
	return a.state == "q2"
}

func main() {
	afd := AFD{}
	fmt.Println(afd.isAccepted("110")) 
	fmt.Println(afd.isAccepted("101")) 
	fmt.Println(afd.isAccepted("1001"))
	fmt.Println(afd.isAccepted("01"))  
	fmt.Println(afd.isAccepted("00"))  
}

// Entendi da questão que a string não precisa conter apenas dois '1's, mas sim que a quantidade de '1's precisa ser igual a 2 
