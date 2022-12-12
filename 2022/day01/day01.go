package day01

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func First() {
	file, _ := os.Open("day01/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	maxValue := 0
	currentCount := 0

	for sc.Scan() {
		currText := sc.Text()
		if currText == "" {
			if currentCount > maxValue {
				maxValue = currentCount
			}
			currentCount = 0

		} else {

			number, err := strconv.Atoi(currText)
			if err != nil {
				panic(err)
			}
			currentCount += number
		}
	}

	fmt.Println(maxValue)
}

func Two() {
	file, _ := os.Open("day01/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	m := []int{}

	count := 0

	for sc.Scan() {
		currText := sc.Text()
		if currText == "" {
			m = append(m, count)
			count = 0
		} else {

			number, err := strconv.Atoi(currText)
			if err != nil {
				panic(err)
			}
			count += number
		}
	}
	m = append(m, count)

	sort.Slice(m, func(i, j int) bool {
		return m[i] > m[j]
	})

	fmt.Println(m[0] + m[1] + m[2])
}
