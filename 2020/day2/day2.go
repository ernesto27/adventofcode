package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func First() {

	var input []string

	file, _ := os.Open("day2/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		input = append(input, sc.Text())
	}

	finalCount := 0
	for _, i := range input {
		slice := strings.Split(i, " ")

		wordToFind := strings.Replace(slice[1], ":", "", -1)

		nums := strings.Split(slice[0], "-")
		min, _ := strconv.Atoi(nums[0])
		max, _ := strconv.Atoi(nums[1])

		password := strings.Split(slice[2], "")

		count := 0
		for _, p := range password {
			if p == wordToFind {
				count++
			}
		}

		if count >= min && count <= max {
			finalCount++
		}
	}

	fmt.Println(finalCount)
}

func Second() {

	var input []string

	file, _ := os.Open("day2/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		input = append(input, sc.Text())
	}

	finalCount := 0
	for _, i := range input {
		slice := strings.Split(i, " ")

		wordToFind := strings.Replace(slice[1], ":", "", -1)

		nums := strings.Split(slice[0], "-")
		min, _ := strconv.Atoi(nums[0])
		max, _ := strconv.Atoi(nums[1])

		password := strings.Split(slice[2], "")

		count := 0
		if password[min-1] == wordToFind {
			count++
		}
		if password[max-1] == wordToFind {
			count++
		}

		if count == 1 {
			finalCount++
		}
	}

	fmt.Println(finalCount)
}
