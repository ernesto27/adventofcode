package day06

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func A() {

	file, _ := ioutil.ReadFile("day06/input.txt")
	input := strings.Split(string(file), "\n")

	var data [][]string
	for _, word := range input {
		row := []string{}
		for _, w := range word {
			row = append(row, string(w))
		}
		data = append(data, row)
	}

	var response string

	for x := 0; x < len(data[0]); x++ {
		m := make(map[string]int)
		for y := 0; y < len(data); y++ {
			m[data[y][x]]++
		}

		fmt.Println(m)
		var char string
		max := 0
		for k, v := range m {
			if v > max {
				max = v
				char = k
			}
		}
		response += char
	}
	fmt.Println(response)
}

func B() {

	file, _ := ioutil.ReadFile("day06/input.txt")
	input := strings.Split(string(file), "\n")

	var data [][]string
	for _, word := range input {
		row := []string{}
		for _, w := range word {
			row = append(row, string(w))
		}
		data = append(data, row)
	}

	var response string

	for x := 0; x < len(data[0]); x++ {
		m := make(map[string]int)
		for y := 0; y < len(data); y++ {
			m[data[y][x]]++
		}

		fmt.Println(m)
		var char string
		min := math.MaxInt
		for k, v := range m {
			if v < min {
				min = v
				char = k
			}
		}
		response += char
	}
	fmt.Println(response)
}
