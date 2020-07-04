package day1

import "fmt"

func Floor(s string) int {
	floor := 0
	for _, c := range s {
		if c == '(' {
			floor += 1
		} else if c == ')' {
			floor -= 1
		} else {
			panic(fmt.Sprintf("Invalid character %v", c))
		}
	}
	return floor
}
