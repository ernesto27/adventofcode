package day02

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
	result := 0

	for sc.Scan() {
		currText := sc.Text()
		items := strings.Split(currText, "\t")

		currMax := 0
		currMin := math.Inf(1)
		for _, i := range items {
			number, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}

			if number > currMax {
				currMax = number
			}

			n := float64(number)
			if n < currMin {
				currMin = n
			}
		}

		diff := currMax - int(currMin)
		result += diff
	}

	fmt.Println(result)
}

func Second() {
	file, err := os.Open("day02/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	sc := bufio.NewScanner(file)

	result := 0

	for sc.Scan() {
		currText := sc.Text()
		items := strings.Split(currText, "\t")
		fmt.Println(items)

		for x := 0; x < len(items); x++ {
			numberOne, err := strconv.Atoi(items[x])
			if err != nil {
				panic(err)
			}

			for y := x + 1; y < len(items); y++ {
				numberTwo, err := strconv.Atoi(items[y])
				if err != nil {
					panic(err)
				}

				first := numberOne
				second := numberTwo
				if numberTwo > numberOne {
					first = numberTwo
					second = numberOne
				}

				if first%second == 0 {
					fmt.Println(numberOne, numberTwo)
					result += first / second
				}
			}
		}
	}

	fmt.Println(result)
}
