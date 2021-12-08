package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	part1(input, false)
	part2(input, true)
}

func part1(input []string, part2 bool) {
	count := 0
	for _, j := range input {
		a := strings.Split(j, "|")
		for _, d := range strings.Split(a[1], " ") {
			if len(d) == 2 || len(d) == 3 || len(d) == 4 || len(d) == 7 {
				count++
			}

		}
	}
	fmt.Println(count)
}

func part2(input []string, part2 bool) {
	ans := 0

	for _, j := range input {
		a := strings.Split(j, "|")
		mem := map[int][]string{}
		resultDigit := map[int]string{}
		mapping := map[string]string{}
		count := map[string]int{}
		for _, d := range strings.Split(a[0], " ") {
			f := SortString(d)
			if _, exists := mem[len(f)]; exists {
				mem[len(f)] = append(mem[len(f)], f)
			} else {
				mem[len(f)] = []string{f}
			}
			switch len(f) {
			case 2:
				resultDigit[1] = f
			case 3:
				resultDigit[7] = f
			case 4:
				resultDigit[4] = f
			case 7:
				resultDigit[8] = f
			}
			for _, c := range f {
				count[string(c)]++
			}
		}
		seven := ""
		for k, v := range count {
			switch v {
			case 4:
				mapping["e"] = k
			case 6:
				mapping["b"] = k
			case 9:
				mapping["f"] = k
			case 7:
				seven += k
			}
		}
		mapping["a"] = SubSet(mem[3][0], mem[2][0])
		mapping["c"] = SubSet(mem[2][0], mapping["f"])
		mapping["d"] = SubSet(SubSet(mem[4][0], mem[2][0]), mapping["b"])
		mapping["g"] = SubSet(seven, mapping["d"])

		resultDigit[0] = SortString(mapping["a"] + mapping["b"] + mapping["c"] + mapping["e"] + mapping["f"] + mapping["g"])
		resultDigit[2] = SortString(mapping["a"] + mapping["c"] + mapping["d"] + mapping["e"] + mapping["g"])
		resultDigit[3] = SortString(mapping["a"] + mapping["c"] + mapping["d"] + mapping["f"] + mapping["g"])
		resultDigit[5] = SortString(mapping["a"] + mapping["b"] + mapping["d"] + mapping["f"] + mapping["g"])
		resultDigit[6] = SortString(mapping["a"] + mapping["b"] + mapping["d"] + mapping["e"] + mapping["f"] + mapping["g"])
		resultDigit[9] = SortString(mapping["a"] + mapping["b"] + mapping["c"] + mapping["d"] + mapping["f"] + mapping["g"])

		fmt.Println(resultDigit)

		result := ""
		for _, v := range strings.Split(a[1], " ") {
			w := SortString(v)
			for km, vm := range resultDigit {
				if vm == w {
					result += strconv.Itoa(km)
				}
			}
		}
		dResult, _ := strconv.Atoi(result)
		ans += dResult
	}
	fmt.Println(ans)

}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func SubSet(main string, sub string) string {
	new := ""
mainLoop:
	for _, c := range main {
		for _, d := range sub {
			if c == d {
				continue mainLoop
			}
		}
		new += string(c)
	}
	return new
}
