package day02

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func First() {
	input := [][]int{}

	file, err := os.Open("day02/day02.txt")
	if err != nil {
		fmt.Println(err, "\nYou need a file called input.txt in this directory")
		return
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		a := strings.Split(sc.Text(), "x")
		var curr []int
		for _, x := range a {
			n, _ := strconv.Atoi(x)
			curr = append(curr, n)
		}
		input = append(input, curr)
	}
	fmt.Println(input)

	resp := 0
	for y := 0; y < len(input); y++ {
		first := input[y][0] * input[y][1]
		second := input[y][1] * input[y][2]
		third := input[y][0] * input[y][2]

		smallArea := first

		if second < smallArea {
			smallArea = second
		}

		if third < smallArea {
			smallArea = third
		}

		squareFeet := (2*first + 2*second + 2*third) + smallArea
		resp += squareFeet
	}

	fmt.Println(resp)
}

func Second() {
	input := [][]int{}

	file, err := os.Open("day02/day02.txt")
	if err != nil {
		fmt.Println(err, "\nYou need a file called input.txt in this directory")
		return
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		a := strings.Split(sc.Text(), "x")
		var curr []int
		for _, x := range a {
			n, _ := strconv.Atoi(x)
			curr = append(curr, n)
		}
		input = append(input, curr)
	}
	fmt.Println(input)

	resp := 0
	for y := 0; y < len(input); y++ {
		sort.Ints(input[y])
		ribbon := input[y][0] + input[y][0] + input[y][1] + input[y][1]
		resp += (input[y][0] * input[y][1] * input[y][2]) + ribbon
	}

	fmt.Println(resp)
}
