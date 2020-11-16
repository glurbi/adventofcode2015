package utils

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func Input(file string) string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

func Assert(t *testing.T, message string, actual int, expected int) {
	if actual != expected {
		t.Errorf("%s, got %v, expected %v", message, actual, expected)
	}
}
