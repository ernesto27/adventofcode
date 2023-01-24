package day03

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func A() {
	input, _ := ioutil.ReadFile("day03/input.txt")

	triangleStrings := strings.Split(string(input), "\n")
	var triangleCount int

	sides := make([]int, 3)

	for _, triangleString := range triangleStrings {
		for i, side := range strings.Fields(triangleString) {
			sides[i], _ = strconv.Atoi(side)
		}
		sort.Ints(sides)

		if sides[0]+sides[1] > sides[2] {
			triangleCount++
		}
	}

	println(triangleCount)
}

func B() {
	var data [][]int
	input, _ := ioutil.ReadFile("day03/input.txt")

	triangleStrings := strings.Split(string(input), "\n")
	var triangleCount int

	for _, triangleString := range triangleStrings {
		sides := make([]int, 3)
		for i, side := range strings.Fields(triangleString) {
			sides[i], _ = strconv.Atoi(side)
		}
		data = append(data, sides)
	}

	for y := 0; y < len(data); y += 3 {
		for x := 0; x < len(data[y]); x++ {
			sides := make([]int, 3)
			sides[0] = data[y][x]
			sides[1] = data[y+1][x]
			sides[2] = data[y+2][x]
			sort.Ints(sides)

			if sides[0]+sides[1] > sides[2] {
				triangleCount++
			}
		}
	}
	fmt.Println(triangleCount)
}
