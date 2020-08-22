package adventofcode2015

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func input(file string) string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

func assert(t *testing.T, message string, actual int, expected int) {
	if actual != expected {
		t.Errorf("%s, got %v, expected %v", message, actual, expected)
	}
}
