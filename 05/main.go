package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var input []string

	for scanner.Scan() {
		v := scanner.Text()
		input = append(input, v)
	}

	file.Close()
	part1(input, false)
	part1(input, true)
}

type point struct {
	x int
	y int
}

func part1(input []string, part2 bool) {
	mem := map[point]int{}

	for _, iv := range input {
		allNumbers := []int{}
		for _, j := range strings.Split(strings.ReplaceAll(iv, " -> ", ","), ",") {
			ij, _ := strconv.Atoi(j)
			allNumbers = append(allNumbers, ij)
		}

		if allNumbers[0] == allNumbers[2] {
			min := allNumbers[1]
			max := allNumbers[3]
			if allNumbers[1] > allNumbers[3] {
				min = allNumbers[3]
				max = allNumbers[1]
			}

			for i := min; i <= max; i++ {
				p := point{
					x: allNumbers[0],
					y: i,
				}

				mem[p]++
			}
		} else if allNumbers[1] == allNumbers[3] {
			min := allNumbers[0]
			max := allNumbers[2]
			if allNumbers[0] > allNumbers[2] {
				min = allNumbers[2]
				max = allNumbers[0]
			}

			for i := min; i <= max; i++ {
				p := point{
					x: i,
					y: allNumbers[1],
				}

				mem[p]++
			}
		} else {
			if !part2 {
				continue
			}

			revertx := false
			x1 := allNumbers[0]
			x2 := allNumbers[2]
			if allNumbers[0] > allNumbers[2] {
				revertx = true
			}

			reverty := false
			y1 := allNumbers[1]
			if allNumbers[1] > allNumbers[3] {
				reverty = true
			}

			steps := x2 - x1
			if revertx {
				steps = x1 - x2
			}

			for i := 0; i <= steps; i++ {
				x := x1 + i
				y := y1 + i

				if revertx {
					x = x1 - i
				}

				if reverty {
					y = y1 - i
				}

				p := point{
					x: x,
					y: y,
				}

				mem[p]++
			}

		}

	}

	count := 0
	for _, v := range mem {
		if v > 1 {
			count++
		}
	}

	fmt.Println(count)
}
