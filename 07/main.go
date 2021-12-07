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
	parts(input, false)
	parts(input, true)
}

func parts(input []string, part2 bool) {
	in := []int{}
	min := -1
	max := -1
	for _, j := range strings.Split(input[0], ",") {
		ij, _ := strconv.Atoi(j)
		in = append(in, ij)
		if ij < min || min == -1 {
			min = ij
		}
		if ij > max || max == -1 {
			max = ij

		}
	}

	fuelMap := []int{}
	lastFuel := 0
	for i := 0; i <= max; i++ {
		lastFuel += i
		fuelMap = append(fuelMap, lastFuel)
	}

	minFuel := -1
	for i := min; i < max; i++ {
		fuel := 0
		for _, val := range in {
			fuelUsed := val - i
			if val-i < 0 {
				fuelUsed = -fuelUsed
			}
			if !part2 {
				fuel += fuelUsed
			} else {
				fuel += fuelMap[fuelUsed]
			}
		}

		if fuel < minFuel || minFuel == -1 {
			minFuel = fuel
		}
	}

	fmt.Println(minFuel)
}
