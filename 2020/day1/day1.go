package day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func First() {
	var input []int

	file, _ := os.Open("day1/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		i, _ := strconv.Atoi(sc.Text())
		input = append(input, i)
	}

	mapInput := make(map[int]bool)

	for _, i := range input {
		mapInput[i] = true
	}

	for _, i := range input {
		half := i - 2020
		abs := math.Abs(float64(half))
		intAbs := int(abs)
		if _, ok := mapInput[intAbs]; ok {
			fmt.Println(i * intAbs)
			break
		}
	}
}

func Second() {
	var input []int

	file, _ := os.Open("day1/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		i, _ := strconv.Atoi(sc.Text())
		input = append(input, i)
	}

	mapInput := make(map[int]bool)

	for _, i := range input {
		mapInput[i] = true
	}

	for _, i := range input {
		half := i - 2020
		abs := math.Abs(float64(half))
		intAbs := int(abs)

		for x, _ := range mapInput {
			if _, ok := mapInput[intAbs]; ok {
				continue
			}

			half := x - intAbs
			abs := math.Abs(float64(half))
			if _, ok := mapInput[int(abs)]; ok {
				fmt.Println(intAbs * x * int(abs))
				return
			}

		}
	}
}
