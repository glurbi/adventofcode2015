package adventofcode2015

import (
	"fmt"
	"testing"
)

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
			return p + 1
		}
	}
	panic(fmt.Sprintf("Not going to basement"))
}

func test_day_1_1(t *testing.T, s string, want int) {
	if got := Day_1_1(s); got != want {
		t.Errorf("%s => %v, want %v", s, got, want)
	}
}

func TestDay_1_1(t *testing.T) {
	test_day_1_1(t, "(())", 0)
	test_day_1_1(t, "()()", 0)
	test_day_1_1(t, "(((", 3)
	test_day_1_1(t, "(()(()(", 3)
	test_day_1_1(t, "())", -1)
	test_day_1_1(t, "))(", -1)
	test_day_1_1(t, ")))", -3)
	test_day_1_1(t, ")())())", -3)
	fmt.Printf("Day 1.1 => %v\n", Day_1_1(input("input/day1.txt")))
}

func test_day_1_2(t *testing.T, s string, want int) {
	if got := Day_1_2(s); got != want {
		t.Errorf("%s => %v, want %v", s, got, want)
	}
}

func TestDay_1_2(t *testing.T) {
	test_day_1_2(t, ")", 1)
	test_day_1_2(t, "()())", 5)
	fmt.Printf("Day 1.2 => %v\n", Day_1_2(input("input/day1.txt")))
}
