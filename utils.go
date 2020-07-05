package adventofcode2015

import "fmt"
import "io/ioutil"

func input(file string) string {
	b, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Print(err)
	}
	return string(b)
}
