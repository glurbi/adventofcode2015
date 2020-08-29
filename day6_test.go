package adventofcode2015

import (
	"fmt"
	"testing"
)

func Day_6_1(s string) int {
	count := 0
	return count
}

func TestDay_6_1(t *testing.T) {
	assert(t, "turn on 0,0 through 999,999", Day_6_1("turn on 0,0 through 999,999"), 1000)
	fmt.Printf("Day 6.1 => %v\n", Day_6_1(input("input/day6.txt")))
}
