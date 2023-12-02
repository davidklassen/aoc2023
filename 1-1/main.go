package main

import (
	"bufio"
	"flag"
	"fmt"

	"hithub.com/davidklassen/aoc2023/aoc"
)

var (
	token = flag.String("token", "", "AoC token")
)

func firstDigit(s string) int {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return int(c - '0')
		}
	}
	return -1
}

func lastDigit(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			return int(s[i] - '0')
		}
	}
	return -1
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
