package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func parts(input []string, part2 bool) {
	mem := map[string]string{}
	init := ""
	for i, j := range input {
		if j == "" {
			continue
		}
		if i == 0 {
			init = j
			continue
		}
		a := strings.Split(j, " -> ")
		mem[a[0]] = a[1]
	}

	currentCount := map[string]int{}
	for j := 0; j < len(init)-1; j++ {
		currentCount[init[j:j+2]]++
	}

	l := 10
	if part2 {
		l = 40
	}

	for i := 0; i < l; i++ {
		nextCount := map[string]int{}
		for k, v := range currentCount {
			r := mem[k]
			nextCount[string(k[0])+r] += v
			nextCount[r+string(k[1])] += v
		}
		currentCount = nextCount
	}
	count := map[byte]int{}
	for k, v := range currentCount {
		count[k[0]] += v
	}
	count[init[len(init)-1]]++

	min := -1
	max := 0
	for _, v := range count {
		if v >= max {
			max = v
		}
		if v <= min || min == -1 {
			min = v
		}
	}

	fmt.Println(max - min)
}
