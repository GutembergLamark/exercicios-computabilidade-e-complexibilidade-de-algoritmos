package main

import "fmt"

type AFN struct {
	startState    string
	acceptStates  map[string]bool
	currentStates map[string]bool
}

func NewAFN() *AFN {
	return &AFN{
		startState:   "q0",
		acceptStates: map[string]bool{"q0": true, "q1": true, "q2": true},
		currentStates: map[string]bool{
			"q0": true,
		},
	}
}

func (afn *AFN) transition(input rune) {
	nextStates := make(map[string]bool)
	for state := range afn.currentStates {
		if state == "q0" {
			if input == '0' {
				nextStates["q1"] = true
			} else if input == '1' {
				nextStates["q2"] = true
			}
		} else if state == "q1" && input == '1' {
			nextStates["q2"] = true
		} else if state == "q2" && input == '0' {
			nextStates["q1"] = true
		}
	}
	afn.currentStates = nextStates
}

func (afn *AFN) IsAccepted(input string) bool {
	afn.currentStates = map[string]bool{
		afn.startState: true,
	}
	for _, c := range input {
		afn.transition(c)
	}

	return len(afn.currentStates) > 0
}

func main() {
	afn := NewAFN()
	fmt.Println(afn.IsAccepted("010101")) 
	fmt.Println(afn.IsAccepted("101010")) 
	fmt.Println(afn.IsAccepted("1101"))   
	fmt.Println(afn.IsAccepted("100"))    
	fmt.Println(afn.IsAccepted(""))       
}
