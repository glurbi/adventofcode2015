package day1

import "fmt"
import "testing"
import "io/ioutil"

func TestFloor(t *testing.T) {
	test(t, "(())", 0)
	test(t, "()()", 0)
	test(t, "(((", 3)
	test(t, "(()(()(", 3)
	test(t, "())", -1)
	test(t, "))(", -1)
	test(t, ")))", -3)
	test(t, ")())())", -3)
}

func TestFloorMain(t *testing.T) {
	b, err := ioutil.ReadFile("day1.txt")
    if err != nil {
        fmt.Print(err)
	}
	fmt.Printf("Day 1 - %v\n", Floor(string(b)))
}

func test(t *testing.T, s string, want int) {
    if got := Floor(s); got != want {
        t.Errorf("Floor(%s) = %v, want %v", s, got, want)
    }
}