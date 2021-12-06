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
	parts(input, 80)
	parts(input, 256)
}

func parts(input []string, len int) {
	mem := [9]int{}
	for _, j := range strings.Split(input[0], ",") {
		ij, _ := strconv.Atoi(j)
		mem[ij]++
	}

	for i := 0; i < len; i++ {
		nextMem := [9]int{}
		for k, v := range mem {
			if k == 0 {
				nextMem[6] += v
				nextMem[8] += v
			} else {
				nextMem[k-1] += v
			}
		}
		mem = nextMem
	}

	count := 0
	for _, v := range mem {
		count += v
	}

	fmt.Println(count)
}
