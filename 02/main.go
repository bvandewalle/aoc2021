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
	fmt.Println("start")

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

	part2(input)
}

func part1(input []string) {
	h := 0
	v := 0
	for _, iv := range input {
		r := strings.Split(iv, " ")
		val, _ := strconv.Atoi(r[1])
		switch r[0] {
		case "up":
			v -= val
		case "down":
			v += val
		case "forward":
			h += val
		}
	}

	fmt.Println(v * h)
}

func part2(input []string) {
	h := 0
	v := 0
	aim := 0
	for _, iv := range input {
		r := strings.Split(iv, " ")
		val, _ := strconv.Atoi(r[1])
		switch r[0] {
		case "up":
			aim -= val
		case "down":
			aim += val
		case "forward":
			h += val
			v += aim * val
		}
	}

	fmt.Println(v * h)
}
