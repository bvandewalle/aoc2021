package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("start")

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var input []int

	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		input = append(input, v)
	}

	file.Close()

	part2(input)
}

func part1(input []int) {
	count := 0
	for i, iv := range input {
		if i > 0 {
			if iv > input[i-1] {
				count++
			}
		}
	}

	fmt.Println(count)
}

func part2(input []int) {
	count := 0
	for i, iv := range input {
		if i > 2 {
			sum1 := input[i-3] + input[i-2] + input[i-1]
			sum2 := input[i-2] + input[i-1] + iv
			if sum2 > sum1 {
				count++
			}
		}
	}

	fmt.Println(count)
}
