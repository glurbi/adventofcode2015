package adventofcode2015

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func Day_2_1(s string) int {
	total := 0
	for _, line := range strings.Split(s, "\n") {
		dims := strings.Split(line, "x")
		x, _ := strconv.Atoi(dims[0])
		y, _ := strconv.Atoi(dims[1])
		z, _ := strconv.Atoi(dims[2])
		lwh := []int{x, y, z}
		sort.Ints(lwh)
		surface := 2*x*y + 2*x*z + 2*y*z + lwh[0]*lwh[1]
		total += surface
	}
	return total
}

func Day_2_2(s string) int {
	total := 0
	for _, line := range strings.Split(s, "\n") {
		dims := strings.Split(line, "x")
		x, _ := strconv.Atoi(dims[0])
		y, _ := strconv.Atoi(dims[1])
		z, _ := strconv.Atoi(dims[2])
		lwh := []int{x, y, z}
		sort.Ints(lwh)
		ribbon := lwh[0]*2 + lwh[1]*2 + x*y*z
		total += ribbon
	}
	return total
}

func test_day_2_1(t *testing.T, s string, want int) {
	if got := Day_2_1(s); got != want {
		t.Errorf("%s => %v, want %v", s, got, want)
	}
}

func TestDay_2_1(t *testing.T) {
	test_day_2_1(t, "2x3x4", 58)
	test_day_2_1(t, "1x1x10", 43)
	test_day_2_1(t, "1x1x10\n2x3x4", 58+43)
	fmt.Printf("Day 2.1 => %v\n", Day_2_1(input("input/day2.txt")))
}

func test_day_2_2(t *testing.T, s string, want int) {
	if got := Day_2_2(s); got != want {
		t.Errorf("%s => %v, want %v", s, got, want)
	}
}

func TestDay_2_2(t *testing.T) {
	test_day_2_2(t, "2x3x4", 34)
	test_day_2_2(t, "1x1x10", 14)
	test_day_2_2(t, "1x1x10\n2x3x4", 34+14)
	fmt.Printf("Day 2.2 => %v\n", Day_2_2(input("input/day2.txt")))
}
