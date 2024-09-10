package main

import (
	"fmt"
)

func check(s1, s2 string) bool {
	return s1 == s2
}

func main() {
	s1 := encode_data("bachlang364", "1234567890", 2)
	s2 := encode_data("bachlang364", "1234567890", 2)
	
	if check(s1, s2) {
		fmt.Println("true")
		fmt.Println(s1)
	} else {
		fmt.Println("false")
	}
	
}

// output: true
// 55+51-03*97+51V31c45P10a48%54Y98-11b97=52Q57V99R08c61(05R54V50(49+52c54+499850975199521045310854975511056103575148545452