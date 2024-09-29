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
			"q0": true,
			"q1": true,
		},
		startState: "q0",
		acceptStates: map[string]bool{
			"q0": true,
			"q1": true,
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
			switch state {
			case "q0":
				if c == 'a' {
					nextStates["q0"] = true
				} else if c == 'b' {
					nextStates["q1"] = true
				}
			case "q1":
				if c == 'b' {
					nextStates["q1"] = true
				}
			}
		}
		currentStates = nextStates
	}

	_, acceptedInQ0 := currentStates["q0"]
	_, acceptedInQ1 := currentStates["q1"]
	return acceptedInQ0 || acceptedInQ1
}

func main() {
	afn := NewAFN()
	fmt.Println(afn.isAccepted("aaa"))   
	fmt.Println(afn.isAccepted("aabbb")) 
	fmt.Println(afn.isAccepted("aaab"))  
	fmt.Println(afn.isAccepted("aabb"))  
	fmt.Println(afn.isAccepted("abba"))  
	fmt.Println(afn.isAccepted("b"))     
	fmt.Println(afn.isAccepted("ba"))    
}
