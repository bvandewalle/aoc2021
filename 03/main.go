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

	part2(input)
}

func part1(input []string) {
	bits := make([]int, len(input[0]))
	count := 0
	for _, iv := range input {
		count++
		for i := 0; i < len(input[0]); i++ {
			if iv[i] == '1' {
				bits[i]++
			}
		}
	}

	gamma := ""
	epsilon := ""
	for i := 0; i < len(input[0]); i++ {
		if bits[i] > count/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	ig, _ := strconv.ParseInt(gamma, 2, 64)
	ie, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println(ig * ie)
}

func part2(input []string) {
	fmt.Println(part2Rec(input, 0, false) * part2Rec(input, 0, true))
}

func part2Rec(input []string, currentBin int, co2 bool) int {
	if len(input) <= 1 {
		o, _ := strconv.ParseInt(input[0], 2, 64)
		return int(o)
	}

	bits := 0
	count := 1
	for _, iv := range input {
		count++
		if iv[currentBin] == '1' {
			bits++
		}
	}

	nextInput := []string{}
	k := co2
	if bits >= count/2 {
		k = !co2
	}

	km := "0"
	if k {
		km = "1"
	}

	for _, iv := range input {
		if iv[currentBin] == km[0] {
			nextInput = append(nextInput, iv)
		}
	}

	return part2Rec(nextInput, currentBin+1, co2)
}
