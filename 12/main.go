package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
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
	mem := map[string][]string{}
	for _, j := range input {
		a := strings.Split(j, "-")
		if _, exists := mem[a[0]]; !exists {
			mem[a[0]] = []string{a[1]}
		} else {
			mem[a[0]] = append(mem[a[0]], a[1])
		}
		if _, exists := mem[a[1]]; !exists {
			mem[a[1]] = []string{a[0]}
		} else {
			mem[a[1]] = append(mem[a[1]], a[0])
		}
	}

	fmt.Println(visitRec(mem, map[string]bool{}, "start", part2))
}

func visitRec(mem map[string][]string, visited map[string]bool, current string, part2 bool) int {
	if current == "end" {
		return 1
	}

	setDouble := false
	if IsLower(current) && current != "start" {
		if _, ex := visited[current]; ex {
			if _, ex2 := visited["double"]; ex2 || !part2 {
				return 0
			} else {
				setDouble = true
			}
		}
	}

	newVisited := copyMap(visited)
	newVisited[current] = true
	if setDouble {
		newVisited["double"] = true
	}

	count := 0
	for _, p := range mem[current] {
		if p == "start" {
			continue
		}
		count += visitRec(mem, newVisited, p, part2)
	}

	return count
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}

	return true
}

func copyMap(visited map[string]bool) map[string]bool {
	targetMap := make(map[string]bool)
	for key, value := range visited {
		targetMap[key] = value
	}

	return targetMap
}
