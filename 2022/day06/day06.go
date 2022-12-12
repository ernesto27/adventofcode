package day06

import (
	"bufio"
	"fmt"
	"os"
)

func First() {
	file, _ := os.Open("day06/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	sc.Scan()
	text := sc.Text()
	l := len(text) - 4

	for x := 0; x <= l; x++ {
		count := 0
		m := make(map[string]bool)
		for j := x; j <= x+3; j++ {
			_, ok := m[string(text[j])]
			if ok {
				break
			}
			count++
			m[string(text[j])] = true
		}

		if count == 4 {
			fmt.Println(x + 4)
			break
		}

		fmt.Println("-----")
	}
}

func Two() {
	file, _ := os.Open("day06/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	sc.Scan()
	text := sc.Text()
	l := len(text) - 14

	for x := 0; x <= l; x++ {
		count := 0
		m := make(map[string]bool)
		for j := x; j <= x+13; j++ {
			_, ok := m[string(text[j])]
			if ok {
				break
			}
			count++
			m[string(text[j])] = true
		}

		if count == 14 {
			fmt.Println(x + 14)
			break
		}

		fmt.Println("-----")
	}
}
