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
		states:       map[string]bool{"q0": true, "q1": true, "q2": true},
		startState:   "q0",
		acceptStates: map[string]bool{"q0": true},
	}
}

func (afd *AFN) IsAccepted(input string) bool {
	currentStates := map[string]bool{afd.startState: true}

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
				if char == '0' {
					nextStates["q2"] = true
				} else {
					nextStates["q1"] = true
				}
			} else if state == "q2" {
				if char == '0' {
					nextStates["q0"] = true
				} else {
					nextStates["q2"] = true
				}
			}
		}
		currentStates = nextStates
	}

	for state := range currentStates {
		if afd.acceptStates[state] {
			return true
		}
	}
	return false
}

func main() {
	afd := NewAFN()
	fmt.Println(afd.IsAccepted("000")) 
	fmt.Println(afd.IsAccepted("001"))  
	fmt.Println(afd.IsAccepted("010"))  
	fmt.Println(afd.IsAccepted("111"))  
	fmt.Println(afd.IsAccepted("0001")) 
	fmt.Println(afd.IsAccepted("0000")) 
}
