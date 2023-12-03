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

func checkGear(schematic [][]byte, pn, i, j int, gears map[[2]int][]int) {
	if schematic[i][j] == '*' {
		gears[[2]int{i, j}] = append(gears[[2]int{i, j}], pn)
	}
}

func findGears(schematic [][]byte) map[[2]int][]int {
	res := make(map[[2]int][]int)

	start := -1
	for i := 0; i < len(schematic); i++ {
		for j := 0; j < len(schematic[i]); j++ {
			if digit(schematic[i][j]) && start == -1 {
				start = j
			}
			if start != -1 && (j == len(schematic[i])-1 || !digit(schematic[i][j+1])) {
				pn, err := strconv.Atoi(string(schematic[i][start : j+1]))
				if err != nil {
					panic(err)
				}

				if start > 0 {
					start--
					checkGear(schematic, pn, i, start, res)
				}
				end := j
				if j < len(schematic[i])-1 {
					end++
					checkGear(schematic, pn, i, end, res)
				}
				if i > 0 {
					for k := start; k <= end; k++ {
						checkGear(schematic, pn, i-1, k, res)
					}
				}
				if i < len(schematic)-1 {
					for k := start; k <= end; k++ {
						checkGear(schematic, pn, i+1, k, res)
					}
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

	for _, v := range findGears(schematic) {
		if len(v) == 2 {
			sum += v[1] * v[0]
		}
	}

	fmt.Println(sum)
}
