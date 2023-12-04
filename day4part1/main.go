package main

import (
	"bufio"
	"flag"
	"fmt"
	"strings"

	"hithub.com/davidklassen/aoc2023/aoc"
)

var (
	token = flag.String("token", "", "AoC token")
)

func main() {
	flag.Parse()

	input, err := aoc.FetchInput(*token, 4)
	if err != nil {
		panic(err)
	}

	var sum int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		card := scanner.Text()
		parts := strings.Split(card, ":")
		parts = strings.Split(parts[1], "|")
		winningList := strings.Split(strings.TrimSpace(parts[0]), " ")
		numbersList := strings.Split(strings.TrimSpace(parts[1]), " ")
		curr := 0
		for _, n := range numbersList {
			for _, w := range winningList {
				if w != "" && n == w {
					if curr == 0 {
						curr = 1
					} else {
						curr = curr * 2
					}
				}
			}
		}
		sum += curr
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
