package main

import (
        "fmt"
        "regexp"
)

func is_accepted(s string) bool {
        re := regexp.MustCompile("^01.*10$")
        return re.MatchString(s)
}

func main() {
        fmt.Println(is_accepted("0110"))  
        fmt.Println(is_accepted("010110")) 
        fmt.Println(is_accepted("10110")) 
        fmt.Println(is_accepted("0111"))  
}