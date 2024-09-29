package main

import "fmt"

type AFN struct {
	states       map[string]bool
	startState   string
	acceptStates map[string]bool
}

func NewAFN() *AFN {
	afn := &AFN{
		states:       make(map[string]bool),
		startState:   "q0",
		acceptStates: make(map[string]bool),
	}
	afn.states["q0"] = true
	afn.states["q1"] = true
	afn.states["q2"] = true
	afn.states["q3"] = true
	afn.acceptStates["q3"] = true
	return afn
}

func (afn *AFN) isAccepted(input string) bool {
	currentStates := map[string]bool{afn.startState: true}

	for _, c := range input {
		nextStates := make(map[string]bool)
		for state := range currentStates {
			switch state {
			case "q0":
				if c == '0' {
					nextStates["q1"] = true
				}
			case "q1":
				if c == '1' {
					nextStates["q2"] = true
				} else if c == '0' {
					nextStates["q1"] = true
				}
			case "q2":
				if c == '0' {
					nextStates["q3"] = true
				} else if c == '1' {
					nextStates["q0"] = true
				}
			case "q3":
				nextStates["q3"] = true
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
	fmt.Println(afn.isAccepted("010"))     
	fmt.Println(afn.isAccepted("0010"))    
	fmt.Println(afn.isAccepted("111"))     
	fmt.Println(afn.isAccepted("0110010")) 
	fmt.Println(afn.isAccepted("101"))     
}
