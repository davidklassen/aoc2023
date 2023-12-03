package main

import (
	"bufio"
	"flag"
	"fmt"
	"strconv"

	"hithub.com/davidklassen/aoc2023/aoc"
)

var (
	token = flag.String("token", "", "AoC token")
)

func digit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isPartNumber(schematic [][]byte, row, start, end int) bool {
	if start > 0 {
		start--
		if schematic[row][start] != '.' {
			return true
		}
	}
	if end < len(schematic[0])-1 {
		end++
		if schematic[row][end] != '.' {
			return true
		}
	}
	if row > 0 {
		for i := start; i <= end; i++ {
			if schematic[row-1][i] != '.' && !digit(schematic[row-1][i]) {
				return true
			}
		}
	}
	if row < len(schematic)-1 {
		for i := start; i <= end; i++ {
			if schematic[row+1][i] != '.' && !digit(schematic[row+1][i]) {
				return true
			}
		}
	}

	return false
}

func partNumbers(schematic [][]byte) []int {
	var res []int
	start := -1
	for i := 0; i < len(schematic); i++ {
		for j := 0; j < len(schematic[i]); j++ {
			if digit(schematic[i][j]) && start == -1 {
				start = j
			}
			if start != -1 && (j == len(schematic[i])-1 || !digit(schematic[i][j+1])) {
				if isPartNumber(schematic, i, start, j) {
					n, err := strconv.Atoi(string(schematic[i][start : j+1]))
					if err != nil {
						panic(err)
					}
					res = append(res, n)
				}
				start = -1
			}
		}
	}

	return res
}

func main() {
	flag.Parse()

	input, err := aoc.FetchInput(*token, 3)
	if err != nil {
		panic(err)
	}

	var sum int
	var schematic [][]byte
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, []byte(line))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for _, n := range partNumbers(schematic) {
		sum += n
	}

	fmt.Println(sum)
}
