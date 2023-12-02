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

func main() {
	flag.Parse()

	input, err := aoc.FetchInput(*token, 2)
	if err != nil {
		panic(err)
	}

	var sum int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		game := scanner.Text()
		parts := strings.Split(game, ":")
		parts = strings.Split(parts[1], ";")
		maxRed, maxGreen, maxBlue := 0, 0, 0
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
						maxRed = n
					}
				case "green":
					if n > maxGreen {
						maxGreen = n
					}
				case "blue":
					if n > maxBlue {
						maxBlue = n
					}
				}
			}
		}
		sum += maxRed * maxGreen * maxBlue
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
