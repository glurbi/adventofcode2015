package day7

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/glurbi/adventofcode2015/utils"
)

func Day_7_1(input string) map[string]uint16 {
	vals := make(map[string]uint16)
	vals["1"] = 1
	for {
		done := true
		for _, line := range strings.Split(input, "\n") {
			ss := strings.Split(line, " -> ")
			operation := ss[0]
			target := ss[1]
			ss = strings.Split(operation, " ")
			if _, ok := vals[target]; ok {
				continue
			}
			if len(ss) == 1 { //assignment
				if v, err := strconv.Atoi(ss[0]); err == nil {
					vals[target] = uint16(v)
					done = false
				} else {
					k := ss[0]
					if v, ok := vals[k]; ok {
						vals[target] = v
						done = false
					}
				}
			} else if ss[0] == "NOT" {
				op := ss[1]
				if v, ok := vals[op]; ok {
					vals[target] = ^v
					done = false
				}
			} else if ss[1] == "AND" {
				op1 := ss[0]
				op2 := ss[2]
				v1, ok1 := vals[op1]
				v2, ok2 := vals[op2]
				if ok1 && ok2 {
					vals[target] = v1 & v2
					done = false
				}
			} else if ss[1] == "OR" {
				op1 := ss[0]
				op2 := ss[2]
				v1, ok1 := vals[op1]
				v2, ok2 := vals[op2]
				if ok1 && ok2 {
					vals[target] = v1 | v2
					done = false
				}
			} else if ss[1] == "LSHIFT" {
				op1 := ss[0]
				v, _ := strconv.Atoi(ss[2])
				v1, ok1 := vals[op1]
				if ok1 {
					vals[target] = v1 << v
					done = false
				}
			} else if ss[1] == "RSHIFT" {
				op1 := ss[0]
				v, _ := strconv.Atoi(ss[2])
				v1, ok1 := vals[op1]
				if ok1 {
					vals[target] = v1 >> v
					done = false
				}
			}
		}
		if done {
			fmt.Printf("%v\n", vals)
			break
		}
	}
	delete(vals, "1")
	return vals
}

func TestDay_7_1(t *testing.T) {
	actual := Day_7_1(`123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`)
	expected := map[string]uint16{
		"d": 72,
		"e": 507,
		"f": 492,
		"g": 114,
		"h": 65412,
		"i": 65079,
		"x": 123,
		"y": 456,
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("got %v, expected %v", actual, expected)
	}
	m := Day_7_1(utils.Input("day7.txt"))
	fmt.Printf("Day 7.1 => %v\na=%v\n", m, m["a"])
}

// 7.2
// replace 19138 by the answer from 7.1 and rerun
