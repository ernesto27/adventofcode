package day05

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func validate(word string) bool {
	valid := false
	vowelCount := 0
	for _, w := range word {
		switch w {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowelCount++
		}
		if vowelCount >= 3 {
			valid = true
			break
		}
	}

	if !valid {
		return false
	}

	valid = false
	// check repeat two letter
	for x := 0; x < len(word); x++ {
		if x < len(word)-1 {
			if word[x] == word[x+1] {
				valid = true
			}
		}
	}

	if !valid {
		return false
	}

	// check invalid
	invalid := []string{"ab", "cd", "pq", "xy"}

	for _, i := range invalid {
		if strings.Contains(word, i) {
			valid = false
			break
		}
	}

	if !valid {
		return false
	}

	return true

}

func First() {
	input, err := ioutil.ReadFile("day05/day05.txt")
	if err != nil {
		panic(err)
	}

	arr := strings.Split(string(input), "\n")

	resp := 0
	for _, a := range arr {
		if validate(a) {
			resp++
		}
	}
	fmt.Println(resp)
}

func Second() {
	input, err := ioutil.ReadFile("day05/day05.txt")
	if err != nil {
		panic(err)
	}

	arr := strings.Split(string(input), "\n")

	resp := 0
	for _, a := range arr {
		if validateB(a) {
			resp++
		}
	}
	fmt.Println(resp)
}

func validateB(word string) bool {
	fmt.Println("VALIDATE FROM ", word)
	valid := false
	for x := 0; x < len(word)-2; x++ {
		if strings.Count(word, word[x:x+2]) >= 2 {
			valid = true
			break
		}

	}

	if !valid {
		return false
	}

	valid = false
	for x := 0; x < len(word)-2; x++ {

		if word[x] == word[x+2] {
			valid = true
			break
		}

	}
	if !valid {
		return false
	}

	return true
}
