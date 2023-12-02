package main

import (
	"bufio"
	"flag"
	"fmt"
	"strings"

	"hithub.com/davidklassen/aoc2023/aoc"
)

var (
	token  = flag.String("token", "", "AoC token")
	digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

func firstDigit(s string) int {
	firstPos := len(s)
	var digit int
	for i, d := range digits {
		if pos := strings.Index(s, d); pos != -1 && pos < firstPos {
			firstPos = pos
			digit = i + 1
		}
	}
	for i, c := range s {
		if c >= '0' && c <= '9' {
			if i < firstPos {
				return int(c - '0')
			} else {
				break
			}
		}
	}
	return digit
}

func lastDigit(s string) int {
	lastPos := -1
	var digit int
	for i, d := range digits {
		if pos := strings.LastIndex(s, d); pos != -1 && pos > lastPos {
			lastPos = pos
			digit = i + 1
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			if i > lastPos {
				return int(s[i] - '0')
			} else {
				break
			}
		}
	}
	return digit
}

func main() {
	flag.Parse()

	input, err := aoc.FetchInput(*token, 1)
	if err != nil {
		panic(err)
	}

	var sum int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		sum += firstDigit(line) * 10
		sum += lastDigit(line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
