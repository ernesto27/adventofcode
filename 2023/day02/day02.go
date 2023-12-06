package day02

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func First() {

	file, err := os.Open("day02/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	index := 1
	sum := 0
	for sc.Scan() {

		t := sc.Text()
		s := strings.Split(t, ":")
		a := strings.Split(s[1], ";")
		isValidGame := true

		for _, i := range a {

			w := strings.Split(i, ",")
			for _, e := range w {
				r := strings.Split(e, " ")
				count, err := strconv.Atoi(strings.TrimSpace(r[1]))
				if err != nil {
					panic(err)
				}

				var blueCount int
				var redCount int
				var greenCount int

				color := strings.TrimSpace(r[2])
				if color == "blue" {
					blueCount += count
				} else if color == "red" {
					redCount += count
				} else if color == "green" {
					greenCount += count
				}

				if redCount > 12 || greenCount > 13 || blueCount > 14 {
					isValidGame = false
					break
				}
			}
		}

		if isValidGame {
			sum += index
		}
		index += 1
	}

	fmt.Println(sum)
}

type round struct {
	red   int
	green int
	blue  int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getNumber(s string) int {
	re := regexp.MustCompile(`\d+`)

	match := re.FindString(s)

	num, err := strconv.Atoi(match)
	check(err)

	return num
}

func isValid(line string) int {
	lineParts := strings.Split(line, ":")
	gameSetsStr := lineParts[1]

	gameSets := strings.Split(gameSetsStr, ";")

	rounds := make([]round, 0)

	for _, set := range gameSets {
		balls := strings.Split(set, ",")

		s := round{}

		for _, ball := range balls {
			n := getNumber(ball)

			if strings.Contains(ball, "red") {
				s.red = n
			} else if strings.Contains(ball, "green") {
				s.green = n
			} else if strings.Contains(ball, "blue") {
				s.blue = n
			}
		}

		rounds = append(rounds, s)
	}

	fewestN := round{}

	for _, round := range rounds {
		if fewestN.red < round.red {
			fewestN.red = round.red
		}

		if fewestN.green < round.green {
			fewestN.green = round.green
		}

		if fewestN.blue < round.blue {
			fewestN.blue = round.blue
		}
	}

	sum := fewestN.red * fewestN.green * fewestN.blue

	return sum
}

func Second() {
	dat, err := os.ReadFile("day02/input.txt")
	check(err)

	lines := strings.Split(string(dat), "\n")

	total := 0

	for _, line := range lines {
		sum := isValid(line)

		total = total + sum
	}

	fmt.Println(total)
}
