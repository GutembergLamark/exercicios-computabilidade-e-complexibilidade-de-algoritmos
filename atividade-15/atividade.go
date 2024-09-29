package main

import "fmt"

type AFN struct {
	startState   string
	acceptStates map[string]bool
	currentState string
}

func NewAFN() *AFN {
	return &AFN{
		startState:   "q0",
		acceptStates: map[string]bool{"q0": true},
		currentState: "q0",
	}
}

func (afn *AFN) transition() {
	if afn.currentState == "q0" {
		afn.currentState = "q1"
	} else if afn.currentState == "q1" {
		afn.currentState = "q0"
	}
}

func (afn *AFN) isAccepted(input string) bool {
	afn.currentState = afn.startState

	for range input {
		afn.transition()
	}

	
	return afn.acceptStates[afn.currentState]
}

func main() {
	afn := NewAFN()
	fmt.Println(afn.isAccepted("ab"))    
	fmt.Println(afn.isAccepted("abab"))  
	fmt.Println(afn.isAccepted("a"))     
	fmt.Println(afn.isAccepted("abc"))   
}