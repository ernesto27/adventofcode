package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func First() {
	// data, err := os.ReadFile("input.txt")
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	path := []string{}
	m := make(map[string][]string)
	firstItem := "AAA"
	endItem := "ZZZ"

	for i, line := range lines {
		if i == 0 {
			for _, l := range line {
				path = append(path, string(l))
			}
			continue
		}
		inputString := line
		if inputString == "" {
			continue
		}

		pattern := regexp.MustCompile(`(\w+)\s*=\s*\((\w+),\s*(\w+)\)`)

		match := pattern.FindStringSubmatch(inputString)

		if match != nil {
			first := match[1]
			second := match[2]
			three := match[3]
			m[first] = []string{second, three}

		} else {
			panic("No match found")
		}
	}

	index := 0
	count := 0
	curr := ""

	for {
		count += 1
		if index > len(path)-1 {
			index = 0
		}

		s := path[index]

		pos := 0
		if s == "R" {
			pos = 1
		}

		if curr == "" {
			curr = firstItem
		}

		cm := m[curr]
		curr = cm[pos]
		index += 1

		if curr == endItem {
			fmt.Println(count)
			fmt.Println("found")
			break
		}
	}
}

type Item struct {
}

func second() {
	data, err := os.ReadFile("simple.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	path := []string{}
	m := make(map[string][]string)

	items := []string{}
	latestZ := make(map[string]bool)

	for i, line := range lines {
		if i == 0 {
			for _, l := range line {
				path = append(path, string(l))
			}
			continue
		}
		inputString := line
		if inputString == "" {
			continue
		}

		pattern := regexp.MustCompile(`(\w+)\s*=\s*\((\w+),\s*(\w+)\)`)

		match := pattern.FindStringSubmatch(inputString)

		if match != nil {
			first := match[1]
			second := match[2]
			three := match[3]
			m[first] = []string{second, three}

			latest := string(first[len(first)-1])

			if latest == "Z" {
				latestZ[first] = true
			} else if latest == "A" {
				items = append(items, first)
			}

		} else {
			panic("No match found")
		}
	}

	index := 0
	count := 0

	for {
		count += 1
		if index > len(path)-1 {
			index = 0
		}

		s := path[index]

		pos := 0
		if s == "R" {
			pos = 1
		}

		isValid := true
		for k, i := range items {
			curr := m[i][pos]
			_, exists := latestZ[curr]

			if !exists {
				isValid = false
			}

			items[k] = curr

		}

		fmt.Println(index)
		if isValid {
			break
		}
		index += 1
	}

	fmt.Println(count)
}

func main() {
	second()
}
