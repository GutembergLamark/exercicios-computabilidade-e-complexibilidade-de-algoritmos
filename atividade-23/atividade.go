package main

import (
	"fmt"
	"regexp"
)

type AFN struct{}

func (AFN) IsAccepted(string string) bool {
	pattern := regexp.MustCompile("101|110") 
	return pattern.MatchString(string)       
}

func main() {
	afn := AFN{}
	fmt.Println(afn.IsAccepted("101"))  
	fmt.Println(afn.IsAccepted("110"))  
	fmt.Println(afn.IsAccepted("0101")) 
	fmt.Println(afn.IsAccepted("0110")) 
	fmt.Println(afn.IsAccepted("1110")) 
	fmt.Println(afn.IsAccepted("000"))  
	fmt.Println(afn.IsAccepted("001"))  
}
