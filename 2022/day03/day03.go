package day03

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func First() {
	file, _ := os.Open("day03/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	result := 0
	for sc.Scan() {
		currText := sc.Text()
		arr := strings.Split(currText, "")
		len := len(arr) / 2

		firstHalf := arr[:len]
		secondHalf := arr[len:]
		repeatLetter := ""

		fmt.Println(arr, firstHalf, secondHalf)

		m := make(map[string]bool)
		for _, first := range firstHalf {
			m[first] = true
		}

		for _, second := range secondHalf {
			if _, ok := m[second]; ok {
				repeatLetter = second
				break
			}
		}
		fmt.Println(repeatLetter)
		letterRune := []rune(repeatLetter)[0]
		ascii := int(letterRune)

		if unicode.IsUpper(letterRune) {
			result += ascii - 38
			fmt.Println(ascii - 38)
		} else {
			result += ascii - 96
			fmt.Println(ascii - 96)
		}
	}

	fmt.Println(result)
}

func Two() {
	file, _ := os.Open("day03/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	m := make(map[string]int)
	count := 0
	result := 0

	for sc.Scan() {
		currText := sc.Text()
		fmt.Println(currText)
		arr := strings.Split(currText, "")

		uniqueValues := make(map[string]bool)
		var uniqueSlice []string
		for _, value := range arr {
			if _, exists := uniqueValues[value]; !exists {
				uniqueValues[value] = true
				uniqueSlice = append(uniqueSlice, value)
			}
		}

		for _, item := range uniqueSlice {
			m[item] += 1
		}

		count += 1

		fmt.Println(count)
		if count == 3 {
			result = updateResult(m, result)
			m = make(map[string]int)
		}

		if count == 6 {
			result = updateResult(m, result)
			count = 0
			m = make(map[string]int)
		}

	}
	fmt.Println("RESULT", result)

}

func updateResult(m map[string]int, result int) int {
	for v := range m {
		if m[v] == 3 {
			letterRune := []rune(v)[0]
			ascii := int(letterRune)

			if unicode.IsUpper(letterRune) {
				result += ascii - 38
			} else {
				result += ascii - 96
			}
		}
	}
	return result
}
