package adventofcode2015

import "strings"
import "strconv"
import "sort"

func Day_2_1(s string) int {
	total := 0
	for _, line := range strings.Split(s, "\n") {
		dims := strings.Split(line, "x")
		x, _ := strconv.Atoi(dims[0])
		y, _ := strconv.Atoi(dims[1])
		z, _ := strconv.Atoi(dims[2])
		lwh := []int { x, y, z }
		sort.Ints(lwh)
		surface := 2*x*y + 2*x*z + 2*y*z + lwh[0] * lwh[1]
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
		lwh := []int { x, y, z }
		sort.Ints(lwh)
		ribbon := lwh[0]*2 + lwh[1]*2 + x*y*z
		total += ribbon
	}
	return total
}
