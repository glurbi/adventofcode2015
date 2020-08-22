package adventofcode2015

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

func Day_4_1(key string) int {
	i := 0
	for {
		data := []byte(fmt.Sprint(key, i))
		hash := md5.Sum(data)
		hashStr := hex.EncodeToString(hash[:])
		if strings.HasPrefix(hashStr, "00000") {
			//fmt.Printf("%x\n", hashStr)
			return i
		}
		i = i + 1
	}
}

func Day_4_2(key string) int {
	i := 0
	for {
		data := []byte(fmt.Sprint(key, i))
		hash := md5.Sum(data)
		hashStr := hex.EncodeToString(hash[:])
		if strings.HasPrefix(hashStr, "000000") {
			//fmt.Printf("%x\n", hashStr)
			return i
		}
		i = i + 1
	}
}

func test_day_4_1(t *testing.T, key string, want int) {
	if got := Day_4_1(key); got != want {
		t.Errorf("%s => %v, want %v", key, got, want)
	}
}

func TestDay_4_1(t *testing.T) {
	test_day_4_1(t, "abcdef", 609043)
	test_day_4_1(t, "pqrstuv", 1048970)
	fmt.Printf("Day 4.1 => %v\n", Day_4_1("ckczppom"))
}

func TestDay_4_2(t *testing.T) {
	fmt.Printf("Day 4.2 => %v\n", Day_4_2("ckczppom"))
}
