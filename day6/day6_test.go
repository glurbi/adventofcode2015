package day6

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/glurbi/adventofcode2015/utils"
)

func Day_6_1(input string) int {
	count := 0
	var grid [1000][1000]bool
	for _, instruction := range strings.Split(input, "\n") {
		split := strings.Split(instruction, " ")
		if split[0] == "toggle" {
			x1, y1 := str2coords(split[1])
			x2, y2 := str2coords(split[3])
			apply(&grid, x1, y1, x2, y2, func(b *bool) { *b = !*b })
		} else if split[1] == "on" { // turn on
			x1, y1 := str2coords(split[2])
			x2, y2 := str2coords(split[4])
			apply(&grid, x1, y1, x2, y2, func(b *bool) { *b = true })
		} else { // turn off
			x1, y1 := str2coords(split[2])
			x2, y2 := str2coords(split[4])
			apply(&grid, x1, y1, x2, y2, func(b *bool) { *b = false })
		}
	}
	for i := 0; i <= 999; i++ {
		for j := 0; j <= 999; j++ {
			if grid[i][j] {
				count++
			}
		}
	}
	return count
}

func str2coords(s string) (int, int) {
	ss := strings.Split(s, ",")
	x, _ := strconv.Atoi(ss[0])
	y, _ := strconv.Atoi(ss[1])
	return x, y
}

func apply(grid *[1000][1000]bool, x1 int, y1 int, x2 int, y2 int, f func(b *bool)) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			f(&grid[i][j])
		}
	}
}

func Day_6_2(input string) int {
	count := 0
	var grid [1000][1000]int
	for _, instruction := range strings.Split(input, "\n") {
		split := strings.Split(instruction, " ")
		if split[0] == "toggle" {
			x1, y1 := str2coords(split[1])
			x2, y2 := str2coords(split[3])
			apply2(&grid, x1, y1, x2, y2, func(b *int) { *b += 2 })
		} else if split[1] == "on" { // turn on
			x1, y1 := str2coords(split[2])
			x2, y2 := str2coords(split[4])
			apply2(&grid, x1, y1, x2, y2, func(b *int) { *b += 1 })
		} else { // turn off
			x1, y1 := str2coords(split[2])
			x2, y2 := str2coords(split[4])
			apply2(&grid, x1, y1, x2, y2, func(b *int) {
				if *b > 0 {
					*b -= 1
				}
			})
		}
	}
	for i := 0; i <= 999; i++ {
		for j := 0; j <= 999; j++ {
			count += grid[i][j]
		}
	}
	return count
}

func apply2(grid *[1000][1000]int, x1 int, y1 int, x2 int, y2 int, f func(b *int)) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			f(&grid[i][j])
		}
	}
}

func TestDay_6_1(t *testing.T) {
	utils.Assert(t, "turn on 0,0 through 999,999", Day_6_1("turn on 0,0 through 999,999"), 1000000)
	utils.Assert(t, "toggle 0,0 through 999,0", Day_6_1("toggle 0,0 through 999,0"), 1000)
	utils.Assert(t, "turn off 499,499 through 500,500", Day_6_1("turn off 499,499 through 500,500"), 0)
	fmt.Printf("Day 6.1 => %v\n", Day_6_1(utils.Input("day6.txt")))
}

func TestDay_6_2(t *testing.T) {
	utils.Assert(t, "turn on 0,0 through 0,0", Day_6_2("turn on 0,0 through 0,0"), 1)
	utils.Assert(t, "toggle 0,0 through 999,999", Day_6_2("toggle 0,0 through 999,999"), 2000000)
	fmt.Printf("Day 6.2 => %v\n", Day_6_2(utils.Input("day6.txt")))
}
