package day04

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func First() {
	file, err := os.Open("day04/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	sc := bufio.NewScanner(file)

	resp := 0

	for sc.Scan() {
		currText := sc.Text()
		items := strings.Split(currText, " ")
		m := make(map[string]bool)
		isValid := true

		for _, i := range items {
			if _, ok := m[i]; ok {
				isValid = false
				break
			}
			m[i] = true
		}

		if isValid {
			resp++
		}
	}

	fmt.Println(resp)
}

func Second() {
	file, err := os.Open("day04/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	sc := bufio.NewScanner(file)

	resp := 0

	for sc.Scan() {
		currText := sc.Text()
		items := strings.Split(currText, " ")
		m := make(map[string]bool)
		isValid := true

		for _, i := range items {
			runes := []rune(i)

			sort.Slice(runes, func(i, j int) bool {
				return runes[i] < runes[j]
			})

			s := string(runes)
			if _, ok := m[s]; ok {
				isValid = false
				break
			}
			m[s] = true
		}

		if isValid {
			resp++
		}
	}

	fmt.Println(resp)
}
