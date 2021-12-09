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

	count := 0
	basins := []int{}
	for i, line := range mem {
		for j, v := range line {
			if i > 0 && mem[i-1][j] <= v {
				continue
			}
			if i < len(mem)-1 && mem[i+1][j] <= v {
				continue
			}
			if j > 0 && mem[i][j-1] <= v {
				continue
			}
			if j < len(mem[0])-1 && mem[i][j+1] <= v {
				continue
			}
			count += (v + 1)
			visited := map[int]bool{}
			basins = append(basins, basinSizeRec(mem, i, j, visited))
		}
	}
	if !part2 {
		fmt.Println(count)
	} else {
		fmt.Println(basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3])
	}
}

func basinSizeRec(mem [][]int, x, y int, state map[int]bool) int {
	if x < 0 || x > len(mem)-1 || y < 0 || y > len(mem[0])-1 || mem[x][y] == 9 {
		return 0
	}
	if _, exists := state[x*len(mem[0])+y]; exists {
		return 0
	}
	state[x*len(mem[0])+y] = true
	return 1 + basinSizeRec(mem, x-1, y, state) + basinSizeRec(mem, x+1, y, state) + basinSizeRec(mem, x, y-1, state) + basinSizeRec(mem, x, y+1, state)
}
