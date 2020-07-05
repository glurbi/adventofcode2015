package day1

import "fmt"

func Floor1(s string) int {
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

func Floor2(s string) int {
	floor := 0
	for p, c := range s {
		if c == '(' {
			floor += 1
		} else if c == ')' {
			floor -= 1
		} else {
			panic(fmt.Sprintf("Invalid character %v", c))
		}
		if floor == -1 {
			return p+1
		}
	}
	panic(fmt.Sprintf("Not going to basement"))
}
