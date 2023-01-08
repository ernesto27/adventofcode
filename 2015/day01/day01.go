package day01

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func First() {
	input, err := os.Open("day01/day01.txt")
	if err != nil {
		fmt.Println(err, "\nYou need a file called input.txt in this directory")
		return
	}
	defer input.Close()
	sc := bufio.NewScanner(input)

	sc.Scan()
	arr := strings.Split(sc.Text(), "")

	res := 0
	for _, c := range arr {
		if c == "(" {
			res++
		} else {
			res--
		}
	}
	fmt.Println(res)
}

func Second() {
	input, err := os.Open("day01/day01.txt")
	if err != nil {
		fmt.Println(err, "\nYou need a file called input.txt in this directory")
		return
	}
	defer input.Close()
	sc := bufio.NewScanner(input)

	sc.Scan()
	arr := strings.Split(sc.Text(), "")

	res := 0
	pos := 0
	for index, c := range arr {
		if c == "(" {
			res++
		} else {
			res--
		}
		if res == -1 {
			pos = index + 1
			break
		}
	}
	fmt.Println(pos)
}
