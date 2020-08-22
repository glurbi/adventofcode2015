package adventofcode2015

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Day_5_1(s string) int {
	count := 0
	for _, line := range strings.Split(s, "\n") {

		// vowels
		vowelCount := 0
		for _, c := range line {
			if strings.ContainsAny(strconv.QuoteRune(c), "aeiou") {
				vowelCount += 1
			}
		}
		if vowelCount < 3 {
			continue
		}

		// twice letter in a row
		twice := false
		for i := range line {
			if i == len(line)-1 {
				break
			}
			if line[i] == line[i+1] {
				twice = true
				break
			}
		}
		if !twice {
			continue
		}

		// does not contain ab, cd, pq, or xy
		found := false
		for _, s := range []string{"ab", "cd", "pq", "xy"} {
			if strings.Contains(line, s) {
				found = true
				break
			}
		}
		if found {
			continue
		}

		// nice string
		count += 1
	}

	return count
}

func test_day_5_1(t *testing.T, s string, want int) {
	if got := Day_5_1(s); got != want {
		t.Errorf("%s => %v, want %v", s, got, want)
	}
}

func TestDay_5_1(t *testing.T) {
	test_day_5_1(t, "ugknbfddgicrmopn", 1)
	test_day_5_1(t, "aaa", 1)
	test_day_5_1(t, "jchzalrnumimnmhp", 0)
	test_day_5_1(t, "haegwjzuvuyypxyu", 0)
	test_day_5_1(t, "dvszwmarrgswjxmb", 0)
	fmt.Printf("Day 5.1 => %v\n", Day_5_1(input("input/day5.txt")))
}
