package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	parts(input, false)
	parts(input, true)

}

func parts(input []string, part2 bool) {
	mem := [][]int{}
	for _, j := range input {
		line := []int{}
		for _, c := range j {
			i, _ := strconv.Atoi(string(c))
			line = append(line, i)
		}
		mem = append(mem, line)
	}

	totalSteps := 999999999999
	if !part2 {
		totalSteps = 100
	}
	flashes := 0
	for step := 0; step < totalSteps; step++ {
		localFlashes := 0
		for i, line := range mem {
			for j, _ := range line {
				mem[i][j] += 1
			}
		}
		for {
			stillFlashing := false
			for i, line := range mem {
				for j, v := range line {
					if v > 9 {
						stillFlashing = true
						flashBoard(mem, i, j)
						flashes++
						localFlashes++
					}
				}
			}
			if !stillFlashing {
				break
			}
		}
		if localFlashes == len(mem)*len(mem[0]) && part2 {
			fmt.Println(step + 1)
			return
		}
	}
	fmt.Println(flashes)
}

func flashBoard(mem [][]int, x, y int) {
	mem[x][y] = 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if x+i < 0 || x+i > len(mem)-1 || y+j < 0 || y+j > len(mem[0])-1 {
				continue
			}
			if mem[x+i][y+j] == 0 {
				continue
			}
			mem[x+i][y+j]++
		}
	}
}
