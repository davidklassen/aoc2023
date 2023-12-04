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

	sum := 0
	id := 1
	res := make(map[int]int)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		card := scanner.Text()
		res[id]++
		sum++
		parts := strings.Split(card, ":")
		parts = strings.Split(parts[1], "|")
		winningList := strings.Split(strings.TrimSpace(parts[0]), " ")
		numbersList := strings.Split(strings.TrimSpace(parts[1]), " ")
		next := id + 1
		for _, n := range numbersList {
			for _, w := range winningList {
				if w != "" && n == w {
					sum += res[id]
					res[next] += res[id]
					next++
				}
			}
		}
		id++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
