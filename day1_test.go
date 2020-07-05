package day1

import "fmt"
import "testing"
import "io/ioutil"

func input() string {
	b, err := ioutil.ReadFile("day1.txt")
    if err != nil {
        fmt.Print(err)
	}
	return string(b)
}

func test1(t *testing.T, s string, want int) {
    if got := Floor1(s); got != want {
        t.Errorf("Floor1(%s) = %v, want %v", s, got, want)
    }
}

func TestFloor1(t *testing.T) {
	test1(t, "(())", 0)
	test1(t, "()()", 0)
	test1(t, "(((", 3)
	test1(t, "(()(()(", 3)
	test1(t, "())", -1)
	test1(t, "))(", -1)
	test1(t, ")))", -3)
	test1(t, ")())())", -3)
}

func TestFloorMain1(t *testing.T) {
	fmt.Printf("Day 1.1 => %v\n", Floor1(input()))
}

func test2(t *testing.T, s string, want int) {
    if got := Floor2(s); got != want {
        t.Errorf("Floor2(%s) = %v, want %v", s, got, want)
    }
}

func TestFloor2(t *testing.T) {
	test2(t, ")", 1)
	test2(t, "()())", 5)
}

func TestFloorMain2(t *testing.T) {
	fmt.Printf("Day 1.2 => %v\n", Floor2(input()))
}

