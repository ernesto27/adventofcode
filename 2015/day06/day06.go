package day06

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type coords struct {
	xStart int
	xEnd   int
	yStart int
	yEnd   int
	value  string
}

func First() {
	input, err := ioutil.ReadFile("day06/day06.txt")
	if err != nil {
		panic(err)
	}

	arr := strings.Split(string(input), "\n")

	var matrix [1000][1000]int
	for _, a := range arr {
		fmt.Println(a)
		b := strings.Split(a, " through ")
		value := ""
		first := []string{}
		cellVal := 0
		if strings.Contains(b[0], "off") {
			value = "off"
			first = strings.Split(strings.Split(b[0], " ")[2], ",")
			cellVal = 0
		} else if strings.Contains(b[0], "on") {
			value = "on"
			first = strings.Split(strings.Split(b[0], " ")[2], ",")
			cellVal = 1
		} else if strings.Contains(b[0], "toggle") {
			value = "toggle"
			first = strings.Split(strings.Split(b[0], " ")[1], ",")
		}

		second := strings.Split(b[1], ",")

		xStart, _ := strconv.Atoi(first[0])
		xEnd, _ := strconv.Atoi(second[0])
		yStart, _ := strconv.Atoi(first[1])
		yEnd, _ := strconv.Atoi(second[1])

		c := coords{
			xStart: xStart,
			xEnd:   xEnd,
			yStart: yStart,
			yEnd:   yEnd,
			value:  value,
		}

		for y := c.yStart; y <= c.yEnd; y++ {
			for x := c.xStart; x <= c.xEnd; x++ {
				if value == "toggle" {
					if matrix[y][x] == 0 {
						cellVal = 1
					} else {
						cellVal = 0
					}
				}
				matrix[y][x] = cellVal
			}
		}
		fmt.Println("-----")

	}

	resp := 0
	for y := 0; y <= 999; y++ {
		for x := 0; x <= 999; x++ {
			if matrix[x][y] == 1 {
				resp++
			}
		}
	}
	fmt.Println(resp)
	return
}

func Second() {
	input, err := ioutil.ReadFile("day06/day06.txt")
	if err != nil {
		panic(err)
	}

	arr := strings.Split(string(input), "\n")

	var matrix [1000][1000]int
	for _, a := range arr {
		fmt.Println(a)
		b := strings.Split(a, " through ")
		value := ""
		first := []string{}
		if strings.Contains(b[0], "off") {
			value = "off"
			first = strings.Split(strings.Split(b[0], " ")[2], ",")
		} else if strings.Contains(b[0], "on") {
			value = "on"
			first = strings.Split(strings.Split(b[0], " ")[2], ",")
		} else if strings.Contains(b[0], "toggle") {
			value = "toggle"
			first = strings.Split(strings.Split(b[0], " ")[1], ",")
		}

		second := strings.Split(b[1], ",")

		xStart, _ := strconv.Atoi(first[0])
		xEnd, _ := strconv.Atoi(second[0])
		yStart, _ := strconv.Atoi(first[1])
		yEnd, _ := strconv.Atoi(second[1])

		c := coords{
			xStart: xStart,
			xEnd:   xEnd,
			yStart: yStart,
			yEnd:   yEnd,
			value:  value,
		}

		for y := c.yStart; y <= c.yEnd; y++ {
			for x := c.xStart; x <= c.xEnd; x++ {
				if value == "on" {
					matrix[y][x] = matrix[y][x] + 1
				} else if value == "off" {
					if matrix[y][x] > 0 {
						matrix[y][x] = matrix[y][x] - 1
					}
				} else if value == "toggle" {
					matrix[y][x] = matrix[y][x] + 2
				}
			}
		}
		fmt.Println("-----")

	}

	resp := 0
	for y := 0; y <= 999; y++ {
		for x := 0; x <= 999; x++ {
			resp += matrix[x][y]
		}
	}
	fmt.Println(resp)
}
