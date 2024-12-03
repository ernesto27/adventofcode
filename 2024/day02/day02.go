package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func First() {
	// data := [][]int{
	// 	{7, 6, 4, 2, 1},
	// 	{1, 2, 7, 8, 9},
	// 	{9, 7, 6, 2, 1},
	// 	{1, 3, 2, 4, 5},
	// 	{8, 6, 4, 4, 1},
	// 	{1, 3, 6, 7, 9},
	// }

	data := [][]int{}
	file, err := os.Open("day02/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)

		items := []int{}

		for _, n := range numbers {
			val, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			items = append(items, val)
		}

		data = append(data, items)
	}

	count := 0
	differMax := 3
	isIncreased := false

	for _, report := range data {
		isValid := true
		for x := 0; x < len(report)-1; x += 1 {
			if report[x] == report[x+1] {
				isValid = false
			}

			if x == 0 {
				if report[x] > report[x+1] {
					isIncreased = false
				} else {
					isIncreased = true
				}
			}

			var diff int
			if !isIncreased {
				diff = report[x] - report[x+1]
			} else {
				diff = report[x+1] - report[x]
			}

			if diff > differMax || diff < 0 {
				isValid = false
				break
			}

		}
		fmt.Println()
		fmt.Println(isValid)
		if isValid {
			count++
		}

	}

	fmt.Println(count)
}
