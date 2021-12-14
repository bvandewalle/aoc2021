package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
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

type point struct {
	x int
	y int
}

func parts(input []string, part2 bool) {
	mem := map[point]bool{}
	folds := [][]int{}
	second := false
	for _, j := range input {
		if j == "" {
			second = true
			continue
		}
		if !second {
			a := strings.Split(j, ",")
			a0, _ := strconv.Atoi(a[0])
			a1, _ := strconv.Atoi(a[1])
			mem[point{a0, a1}] = true
		} else {
			a := strings.Split(j, " ")
			b := strings.Split(a[2], "=")
			b1, _ := strconv.Atoi(b[1])
			b0 := 0
			if b[0] == "y" {
				b0 = 1
			}
			folds = append(folds, []int{b0, b1})
		}
	}

	previousMap := mem
	for i, f := range folds {

		nextMap := map[point]bool{}
		for k, _ := range previousMap {
			if f[0] == 0 {
				if k.x > f[1] {
					newK := point{x: 2*f[1] - k.x, y: k.y}
					nextMap[newK] = true
				} else {
					nextMap[k] = true
				}
			} else {
				if k.y > f[1] {
					newK := point{x: k.x, y: 2*f[1] - k.y}
					nextMap[newK] = true
				} else {
					nextMap[k] = true
				}
			}
		}

		previousMap = nextMap
		if i == 0 && !part2 {
			fmt.Println(len(nextMap))
			return
		}
	}

	upLeft := image.Point{0, 0}
	lowRight := image.Point{50, 50}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	for k := range previousMap {
		img.Set(k.x, k.y, color.White)
	}
	f, _ := os.Create("out.png")
	png.Encode(f, img)
}

func copyMap(visited map[int]int) map[int]int {
	targetMap := make(map[int]int)
	for key, value := range visited {
		targetMap[key] = value
	}

	return targetMap
}
