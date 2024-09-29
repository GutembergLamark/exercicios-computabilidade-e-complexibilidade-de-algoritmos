package main

import (
	"fmt"
)

type AFN struct {
	states       map[string]struct{}
	startState   string
	acceptStates map[string]struct{}
}

func NewAFN() *AFN {
	return &AFN{
		states:       map[string]struct{}{"q0": {}, "q1": {}, "q2": {}},
		startState:   "q0",
		acceptStates: map[string]struct{}{"q2": {}},
	}
}

func (afn *AFN) isAccepted(input string) bool {
	currentStates := map[string]struct{}{afn.startState: {}}

	for _, char := range input {
		nextStates := make(map[string]struct{})
		for state := range currentStates {
			if state == "q0" {
				if char == '0' {
					nextStates["q1"] = struct{}{}
				}
			} else if state == "q1" {
				if char == '1' {
					nextStates["q2"] = struct{}{}
				} else if char == '0' {
					nextStates["q1"] = struct{}{}
				}
			} else if state == "q2" {
				nextStates["q2"] = struct{}{}
			}
		}
		currentStates = nextStates
	}

	for state := range currentStates {
		if _, exists := afn.acceptStates[state]; exists {
			return true
		}
	}
	return false
}

func main() {
	afn := NewAFN()
	fmt.Println(afn.isAccepted("010"))    
	fmt.Println(afn.isAccepted("0011"))   
	fmt.Println(afn.isAccepted("111"))    
	fmt.Println(afn.isAccepted("000"))    
	fmt.Println(afn.isAccepted("000111")) 
}
