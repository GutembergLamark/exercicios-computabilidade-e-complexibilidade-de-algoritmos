package main

import (
	"fmt"
)

type AFN struct {
	states       map[string]bool
	startState   string
	acceptStates map[string]bool
}

func NewAFN() *AFN {
	return &AFN{
		states:       map[string]bool{"q0": true, "q1": true},
		startState:   "q0",
		acceptStates: map[string]bool{"q1": true},
	}
}

func (afn *AFN) isAccepted(input string) bool {
	currentStates := map[string]bool{afn.startState: true}

	for _, char := range input {
		nextStates := make(map[string]bool)
		for state := range currentStates {
			if state == "q0" {
				if char == '0' {
					nextStates["q1"] = true 
				} else {
					nextStates["q0"] = true 
				}
			} else if state == "q1" {
				nextStates["q1"] = true 
			}
		}
		currentStates = nextStates
	}

	for state := range currentStates {
		if afn.acceptStates[state] {
			return true
		}
	}
	return false
}

func main() {
	afn := NewAFN()
	fmt.Println(afn.isAccepted("111"))
	fmt.Println(afn.isAccepted("110"))
	fmt.Println(afn.isAccepted("001"))
	fmt.Println(afn.isAccepted("1"))  
	fmt.Println(afn.isAccepted("0"))  
}
