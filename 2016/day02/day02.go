package day02

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func A() {
	input := []string{}

	file, err := ioutil.ReadFile("day02/input.txt")
	if err != nil {
		panic(err)
	}

	for _, i := range strings.Split(string(file), "\n") {
		input = append(input, i)
	}

	start := 5
	resp := []int{}

	for _, i := range input {
		for _, j := range strings.Split(i, "") {
			fmt.Print(j)
			switch j {
			case "U":
				if start > 3 {
					start -= 3
				}
			case "D":
				if start < 7 {
					start += 3
				}
			case "L":
				if start%3 != 1 {
					start--
				}
			case "R":
				if start%3 != 0 {
					start++
				}
			}
		}
		resp = append(resp, start)
		fmt.Println()
	}

	fmt.Println(start)
	fmt.Println(resp)
}

func B() {
	input := []string{}

	file, err := ioutil.ReadFile("day02/input.txt")
	if err != nil {
		panic(err)
	}

	for _, i := range strings.Split(string(file), "\n") {
		input = append(input, i)
	}
	matrix := [][]int{
		{-1, -1, 1, -1, 1},
		{-1, 2, 3, 4, -1},
		{5, 6, 7, 8, 9},
		{-1, 11, 22, 33, -1},
		{-1, -1, 44, -1, -1},
	}

	y := 2
	x := 0
	resp := []int{}

	for _, i := range input {
		for _, j := range strings.Split(i, "") {
			switch j {
			case "U":
				if y != 0 && matrix[y-1][x] != -1 {
					y--
				}
			case "D":
				if y != len(matrix)-1 && matrix[y+1][x] != -1 {
					y++
				}
			case "L":
				if x != 0 && matrix[y][x-1] != -1 {
					x--
				}
			case "R":
				if x != len(matrix[0])-1 && matrix[y][x+1] != -1 {
					x++
				}
			}
		}

		resp = append(resp, matrix[y][x])
	}

	fmt.Println(resp)
}
