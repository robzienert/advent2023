package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

const (
	minWordLength, maxWordLength = 3, 5
)

var (
	//go:embed input.txt
	input string

	digits = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
)

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
	var cur string
	for _, c := range []byte(in) {
		if '0' <= c && c <= '9' {
			v := string(c)
			nums = append(nums, mustInt(v))
			cur = ""
			continue
		}

		cur += string(c)
		match, reset := isEnglishNumber(cur)
		if match > 0 {
			nums = append(nums, match)
			cur = string(c)
		}
		if reset {
			cur = string(c)
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

// first ret: is match?
// second ret: reset?
func isEnglishNumber(str string) (int, bool) {
	if len(str) < minWordLength {
		return 0, false
	}

	digit := matchDigit(str)
	return digit, digit > 0
}

func matchDigit(str string) int {
	for name, digit := range digits {
		if strings.Contains(str, name) {
			return digit
		}
	}
	return 0
}
