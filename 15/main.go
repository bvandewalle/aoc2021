package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")

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

type point struct {
	x int
	y int
}

type funcHolder struct {
	value int
	x     int
	y     int
}

func parts(input []string, part2 bool) {
	mem := [][]int{}
	bestPath := map[point]int{}
	for _, j := range input {
		line := []int{}
		for _, a := range j {
			a0, _ := strconv.Atoi(string(a))
			line = append(line, a0)
		}
		mem = append(mem, line)
	}

	if part2 {
		mem2 := [][]int{}
		for i := 0; i < 5; i++ {
			for ii, l := range mem {
				line := []int{}
				for j := 0; j < 5; j++ {
					for jj, _ := range l {
						incr := i + j
						val := mem[ii][jj] + incr
						if val >= 10 {
							val -= 9
						}
						line = append(line, val)
					}
				}
				mem2 = append(mem2, line)
			}
		}
		mem = mem2
	}

	pathIter(mem, bestPath)
	fmt.Println(bestPath[point{x: len(mem[0]) - 1, y: len(mem) - 1}])
}

func pathIter(mem [][]int, bestPath map[point]int) {
	funcHolders := []funcHolder{funcHolder{value: 0, x: 0, y: 0}}

	for len(funcHolders) > 0 {
		next := funcHolders[0]
		funcHolders = funcHolders[1:]
		currentPoint := point{x: next.x, y: next.y}
		//fmt.Println(next)
		//fmt.Println(funcHolders)

		if v, ex := bestPath[currentPoint]; ex {
			if next.value >= v {
				continue
			}
		}
		bestPath[currentPoint] = next.value

		if currentPoint.x+1 < len(mem[0]) {
			funcHolders = append(funcHolders, funcHolder{value: next.value + mem[currentPoint.y][currentPoint.x+1], x: currentPoint.x + 1, y: currentPoint.y})
		}
		if currentPoint.x-1 > 0 {
			funcHolders = append(funcHolders, funcHolder{value: next.value + mem[currentPoint.y][currentPoint.x-1], x: currentPoint.x - 1, y: currentPoint.y})
		}
		if currentPoint.y+1 < len(mem) {
			funcHolders = append(funcHolders, funcHolder{value: next.value + mem[currentPoint.y+1][currentPoint.x], x: currentPoint.x, y: currentPoint.y + 1})
		}
		if currentPoint.y-1 > 0 {
			funcHolders = append(funcHolders, funcHolder{value: next.value + mem[currentPoint.y-1][currentPoint.x], x: currentPoint.x, y: currentPoint.y - 1})
		}

		sort.Slice(funcHolders, func(i, j int) bool {
			return funcHolders[i].value < funcHolders[j].value
		})
	}
}
