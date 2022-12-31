package day09

import (
	"bufio"
	"fmt"
	"os"
)

func First() {
	file, err := os.Open("day09/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	sc := bufio.NewScanner(file)

	var text string
	for sc.Scan() {
		text = sc.Text()
	}

	currentScore := 0
	score := 0
	garbage := false

	for i := 0; i < len(text); i++ {
		switch text[i] {
		case '{':
			if !garbage {
				currentScore += 1
			}
		case '}':
			if !garbage {
				score += currentScore
				currentScore -= 1
			}
		case '<':
			if !garbage {
				garbage = true
			}
		case '>':
			garbage = false
		case '!':
			i++
		}
	}
	fmt.Println(score)
}

func Second() {
	file, err := os.Open("day09/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	sc := bufio.NewScanner(file)

	var text string
	for sc.Scan() {
		text = sc.Text()
	}

	score := 0
	garbage := false
	for i := 0; i < len(text); i++ {
		switch c := text[i]; c {
		case '<':
			if garbage {
				score++
			} else {
				garbage = true
			}
		case '>':
			garbage = false
		case '!':
			i++
		default:
			if garbage {
				score++
			}
		}
	}
	fmt.Println(score)
}
