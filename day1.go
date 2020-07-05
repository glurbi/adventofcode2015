package adventofcode2015

import "fmt"

func Day_1_1(s string) int {
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

func Day_1_2(s string) int {
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
