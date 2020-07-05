package adventofcode2015

import "fmt"
import "testing"

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
	fmt.Printf("Day 1.1 => %v\n", Day_1_1(input("day1.txt")))
}

func test_day_1_2(t *testing.T, s string, want int) {
    if got := Day_1_2(s); got != want {
        t.Errorf("%s => %v, want %v", s, got, want)
    }
}

func TestDay_1_2(t *testing.T) {
	test_day_1_2(t, ")", 1)
	test_day_1_2(t, "()())", 5)
	fmt.Printf("Day 1.2 => %v\n", Day_1_2(input("day1.txt")))
}

