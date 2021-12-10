package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	scores := []int{}
	count := 0
line:
	for _, j := range input {
		mem := []rune{}
		for _, c := range j {
			fmt.Print(string(c))
			switch c {
			case '{':
				mem = append(mem, c)
			case '(':
				mem = append(mem, c)
			case '[':
				mem = append(mem, c)
			case '<':
				mem = append(mem, c)
			case '}':
				if mem[len(mem)-1] == '{' {
					mem = mem[:len(mem)-1]
				} else {
					count += 1197
					continue line
				}
			case ')':
				if mem[len(mem)-1] == '(' {
					mem = mem[:len(mem)-1]
				} else {
					count += 3
					continue line
				}
			case ']':
				if mem[len(mem)-1] == '[' {
					mem = mem[:len(mem)-1]
				} else {
					count += 57
					continue line
				}
			case '>':
				if mem[len(mem)-1] == '<' {
					mem = mem[:len(mem)-1]
				} else {
					count += 25137
					continue line
				}
			}
		}
		fmt.Println()
		score := 0
		for i := len(mem) - 1; i >= 0; i-- {
			c := mem[i]

			score *= 5
			switch c {
			case '{':
				score += 3
			case '(':
				score += 1
			case '[':
				score += 2
			case '<':
				score += 4
			}
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	if part2 {
		fmt.Println(scores[(len(scores) / 2)])
	} else {
		fmt.Println(count)
	}
}
