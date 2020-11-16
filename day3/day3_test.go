package day3

import (
	"fmt"
	"testing"

	"github.com/glurbi/adventofcode2015/utils"
)

type location struct {
	x int
	y int
}

func Day_3_1(s string) int {
	x := 0
	y := 0
	m := make(map[location]int)
	m[location{x, y}] += 1
	for _, c := range s {
		if c == '>' {
			x += 1
		} else if c == '<' {
			x -= 1
		} else if c == '^' {
			y += 1
		} else if c == 'v' {
			y -= 1
		} else {
			panic(fmt.Sprintf("Invalid character %v", c))
		}
		m[location{x, y}] += 1
	}
	return len(m)
}

func Day_3_2(s string) int {
	x := [2]int{0, 0}
	y := [2]int{0, 0}
	m := make(map[location]int)
	m[location{0, 0}] += 2
	for i, c := range s {
		o := i % 2
		if c == '>' {
			x[o] += 1
		} else if c == '<' {
			x[o] -= 1
		} else if c == '^' {
			y[o] += 1
		} else if c == 'v' {
			y[o] -= 1
		} else {
			panic(fmt.Sprintf("Invalid character %v", c))
		}
		m[location{x[o], y[o]}] += 1
	}
	return len(m)
}

func test_day_3_1(t *testing.T, s string, want int) {
	if got := Day_3_1(s); got != want {
		t.Errorf("%s => %v, want %v", s, got, want)
	}
}

func TestDay_3_1(t *testing.T) {
	test_day_3_1(t, ">", 2)
	test_day_3_1(t, "^>v<", 4)
	test_day_3_1(t, "^v^v^v^v^v", 2)
	fmt.Printf("Day 3.1 => %v\n", Day_3_1(utils.Input("input/day3.txt")))
}

func test_day_3_2(t *testing.T, s string, want int) {
	if got := Day_3_2(s); got != want {
		t.Errorf("%s => %v, want %v", s, got, want)
	}
}

func TestDay_3_2(t *testing.T) {
	test_day_3_2(t, "^v", 3)
	test_day_3_2(t, "^>v<", 3)
	test_day_3_2(t, "^v^v^v^v^v", 11)
	fmt.Printf("Day 3.2 => %v\n", Day_3_2(utils.Input("input/day3.txt")))
}
