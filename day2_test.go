package adventofcode2015

import "testing"
import "fmt"

func test_day_2_1(t *testing.T, s string, want int) {
    if got := Day_2_1(s); got != want {
        t.Errorf("%s => %v, want %v", s, got, want)
    }
}

func TestDay_2_1(t *testing.T) {
	test_day_2_1(t, "2x3x4", 58)
	test_day_2_1(t, "1x1x10", 43)
	test_day_2_1(t, "1x1x10\n2x3x4", 58+43)
	fmt.Printf("Day 2.1 => %v\n", Day_2_1(input("day2.txt")))
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
	fmt.Printf("Day 2.2 => %v\n", Day_2_2(input("day2.txt")))
}

