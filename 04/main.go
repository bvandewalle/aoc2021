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
	part1(input)
	part2(input)
}

func part1(input []string) {
	allNumbers := []int{}
	for _, j := range strings.Split(input[0], ",") {
		ij, _ := strconv.Atoi(j)
		allNumbers = append(allNumbers, ij)
	}

	allBoards := [][][]int{}

	currentBoard := [][]int{}
	for i, iv := range input[2:] {
		if (i+1)%6 == 0 && i > 0 {
			allBoards = append(allBoards, currentBoard)
			currentBoard = [][]int{}
			continue
		}

		ii := strings.ReplaceAll(strings.TrimSpace(iv), "  ", " ")

		iline := []int{}
		for _, j := range strings.Split(ii, " ") {
			ij, _ := strconv.Atoi(j)
			iline = append(iline, ij)
		}
		currentBoard = append(currentBoard, iline)
	}
	allBoards = append(allBoards, currentBoard)

	for _, n := range allNumbers {
		for _, b := range allBoards {
			for _, l := range b {
				for i, iv := range l {
					if iv == n {
						l[i] = -1
					}
				}
			}
		}

		for _, b := range allBoards {
		outer1:
			for _, l := range b {
				for _, iv := range l {
					if iv != -1 {
						continue outer1
					}
				}

				fmt.Println(calculateResult(b, n))
				return
			}
		outer2:
			for i := 0; i < len(b[0]); i++ {
				for j := 0; j < len(b); j++ {
					if b[j][i] != -1 {
						continue outer2
					}
				}

				fmt.Println(calculateResult(b, n))
				return
			}

		}
	}
}

func part2(input []string) {
	allNumbers := []int{}
	for _, j := range strings.Split(input[0], ",") {
		ij, _ := strconv.Atoi(j)
		allNumbers = append(allNumbers, ij)
	}
	allBoards := [][][]int{}

	currentBoard := [][]int{}
	for i, iv := range input[2:] {
		if (i+1)%6 == 0 && i > 0 {
			allBoards = append(allBoards, currentBoard)
			currentBoard = [][]int{}
			continue
		}

		ii := strings.ReplaceAll(strings.TrimSpace(iv), "  ", " ")

		iline := []int{}
		for _, j := range strings.Split(ii, " ") {
			ij, _ := strconv.Atoi(j)
			iline = append(iline, ij)
		}
		currentBoard = append(currentBoard, iline)
	}
	allBoards = append(allBoards, currentBoard)

	mem := map[int]bool{}

	for _, n := range allNumbers {

		for _, b := range allBoards {
			for _, l := range b {
				for i, iv := range l {
					if iv == n {
						l[i] = -1
					}
				}
			}
		}

		for in, b := range allBoards {
		outer1:
			for _, l := range b {
				for _, iv := range l {
					if iv != -1 {
						continue outer1
					}
				}

				mem[in] = true
				if len(mem) == len(allBoards) {
					fmt.Println(calculateResult(b, n))
					return
				}
			}
		outer2:
			for i := 0; i < len(b[0]); i++ {
				for j := 0; j < len(b); j++ {
					if b[j][i] != -1 {
						continue outer2
					}
				}

				mem[in] = true
				if len(mem) == len(allBoards) {
					fmt.Println(calculateResult(b, n))
					return
				}
			}
		}
	}
}

func calculateResult(b [][]int, n int) int {
	sum := 0
	for _, l := range b {
		for _, iv := range l {
			if iv != -1 {
				sum += iv
			}
		}
	}

	return sum * n
}
