package main

import (
	"fmt"
	"regexp"
)

func isAccepted(s string) bool {
	re := regexp.MustCompile("110")
	return re.MatchString(s)
}

func main() {
	fmt.Println(isAccepted("101101")) 
	fmt.Println(isAccepted("011"))    
	fmt.Println(isAccepted("111110")) 
}
