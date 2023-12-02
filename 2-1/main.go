package main

import (
	"bufio"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"hithub.com/davidklassen/aoc2023/aoc"
)

var (
	token = flag.String("token", "", "AoC token")
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func validGame(game string) bool {
	parts := strings.Split(game, ":")
	parts = strings.Split(parts[1], ";")

	for _, p := range parts {
		for _, color := range strings.Split(p, ",") {
			pair := strings.Split(strings.TrimSpace(color), " ")
			n, err := strconv.Atoi(pair[0])
			if err != nil {
				panic(err)
			}
			switch pair[1] {
			case "red":
				if n > maxRed {
					return false
				}
			case "green":
				if n > maxGreen {
					return false
				}
			case "blue":
				if n > maxBlue {
					return false
				}
			}
		}
	}

	return true
}

func main() {
	flag.Parse()

	input, err := aoc.FetchInput(*token, 2)
	if err != nil {
		panic(err)
	}

	var sum int
	id := 1
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if validGame(scanner.Text()) {
			sum += id
		}
		id++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
