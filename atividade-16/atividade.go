package main

import "fmt"

type AFD struct {
	state int
}

func NewAFD() *AFD {
	return &AFD{state: 0}
}

func (afd *AFD) IsAccepted(string string) bool {
	afd.state = 0
	for _, c := range string {
		if afd.state == 0 {
			if c == '1' {
				afd.state = 0
			} else if c == '0' {
				afd.state = 1
			}
		} else if afd.state == 1 {
			if c == '1' {
				afd.state = 0
			} else if c == '0' {
				afd.state = 2
			}
		} else if afd.state == 2 {
			if c == '0' {
				afd.state = 2
			} else if c == '1' {
				
			}
		}
	}
	return afd.state == 2
}

func main() {
	afd := NewAFD()
	fmt.Println(afd.IsAccepted("100011")) 
	fmt.Println(afd.IsAccepted("10101"))  
	fmt.Println(afd.IsAccepted("000"))    
	fmt.Println(afd.IsAccepted("000110")) 
	fmt.Println(afd.IsAccepted("1"))      
	fmt.Println(afd.IsAccepted("100"))    
}
