package day07

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func isValidIP(brackets []string, sentences []string) bool {
	for _, w := range brackets {
		run := []rune(w)

		for x := 0; x < len(run)-3; x++ {
			if w[x] == w[x+3] && w[x+1] != w[x] && w[x+1] == w[x+2] {
				return false
			}
		}
	}

	for _, w := range sentences {
		run := []rune(w)
		for x := 0; x < len(run)-3; x++ {
			if w[x] == w[x+3] && w[x+1] != w[x] && w[x+1] == w[x+2] {
				return true
			}
		}
	}

	return false
}

func A() {
	file, _ := ioutil.ReadFile("day07/input.txt")
	input := strings.Split(string(file), "\n")

	count := 0
	r := regexp.MustCompile(`\[([a-z]+)\]`)
	for _, word := range input {
		fmt.Println(word)
		brackets := r.FindAllString(word, -1)
		for i, _ := range brackets {
			brackets[i] = strings.Replace(brackets[i], "[", "", -1)
			brackets[i] = strings.Replace(brackets[i], "]", "", -1)
		}

		s := r.ReplaceAllString(word, " ")
		sentences := strings.Split(s, " ")

		if isValidIP(brackets, sentences) {
			count++
		}
	}

	fmt.Println(count)
}

func isValidIPB(brackets []string, sentences []string) bool {
	m := make(map[string]bool)
	for _, w := range sentences {
		run := []rune(w)
		for x := 0; x < len(run)-2; x++ {
			if w[x] == w[x+2] && w[x] != w[x+1] {
				m[string(w[x+1])+string(w[x])+string(w[x+1])] = true
			}
		}
	}

	for _, w := range brackets {
		run := []rune(w)

		for x := 0; x < len(run)-2; x++ {
			if w[x] == w[x+2] && w[x] != w[x+1] {
				_, ok := m[string(w[x])+string(w[x+1])+string(w[x])]
				if ok {
					return true
				}
			}
		}
	}

	return false
}

func B() {

	file, _ := ioutil.ReadFile("day07/input.txt")
	input := strings.Split(string(file), "\n")

	count := 0
	r := regexp.MustCompile(`\[([a-z]+)\]`)
	for _, word := range input {
		fmt.Println(word)
		brackets := r.FindAllString(word, -1)
		for i, _ := range brackets {
			brackets[i] = strings.Replace(brackets[i], "[", "", -1)
			brackets[i] = strings.Replace(brackets[i], "]", "", -1)
		}

		s := r.ReplaceAllString(word, " ")
		sentences := strings.Split(s, " ")

		if isValidIPB(brackets, sentences) {
			count++
		}
	}

	fmt.Println(count)
}
