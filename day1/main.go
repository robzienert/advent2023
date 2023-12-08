package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var sum int
	for _, ln := range strings.Split(input, "\n") {
		if ln == "" {
			continue
		}
		sum += findValue(ln)
	}

	fmt.Printf("calibration value: %d\n", sum)
}

func findValue(in string) int {
	var nums []int
	for _, c := range []byte(in) {
		if '0' <= c && c <= '9' {
			nums = append(nums, mustInt(string(c)))
			continue
		}
	}

	return mustInt(fmt.Sprintf("%d%d", nums[0], nums[len(nums)-1]))
}

func mustInt(in string) int {
	i, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return i
}
