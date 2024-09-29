package main

import "fmt"

type AFN struct {
	states       map[string]bool
	startState   string
	acceptStates map[string]bool
}

func NewAFN() *AFN {
	return &AFN{
		states: map[string]bool{
			"q0":       true,
			"q1":       true,
			"q_reject": true,
		},
		startState: "q0",
		acceptStates: map[string]bool{
			"q0": true,
		},
	}
}

func (afn *AFN) isAccepted(input string) bool {
	currentStates := map[string]bool{
		afn.startState: true,
	}

	for _, c := range input {
		nextStates := make(map[string]bool)
		for state := range currentStates {
			if state == "q0" {
				if c == 'a' {
					nextStates["q0"] = true
				} else if c == 'b' {
					nextStates["q1"] = true
				}
			} else if state == "q1" {
				if c == 'a' {
					nextStates["q0"] = true
				} else if c == 'b' {
					nextStates["q_reject"] = true
				}
			} else if state == "q_reject" {
				nextStates["q_reject"] = true
			}
		}
		currentStates = nextStates
	}

	return currentStates["q0"]
}

func main() {
	afn := NewAFN()
	fmt.Println(afn.isAccepted("a"))     
	fmt.Println(afn.isAccepted("ab"))    
	fmt.Println(afn.isAccepted("abb"))   
	fmt.Println(afn.isAccepted("ababa")) 
	fmt.Println(afn.isAccepted("bba"))   
	fmt.Println(afn.isAccepted("ba"))    
	fmt.Println(afn.isAccepted("b"))     
}
