package day05

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func First() {
	file, err := os.Open("day05/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	sc := bufio.NewScanner(file)

	nums := []int{}
	for sc.Scan() {
		currText := sc.Text()
		n, err := strconv.Atoi(currText)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}

	currIndex := 0
	steps := 0

	for {
		if currIndex > len(nums)-1 {
			break
		}

		steps += 1
		item := nums[currIndex]
		if item == 0 {
			nums[currIndex] = 1
			continue
		}

		nums[currIndex] += 1
		currIndex += item
	}

	fmt.Println(nums)
	fmt.Println(steps)
}

func Second() {
	file, err := os.Open("day05/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	sc := bufio.NewScanner(file)

	nums := []int{}
	for sc.Scan() {
		currText := sc.Text()
		n, err := strconv.Atoi(currText)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}

	currIndex := 0
	steps := 0

	for {
		if currIndex > len(nums)-1 {
			break
		}

		steps += 1
		item := nums[currIndex]
		if item == 0 {
			nums[currIndex] = 1
			continue
		}

		if item >= 3 {
			nums[currIndex] -= 1
		} else {
			nums[currIndex] += 1
		}

		currIndex += item
	}

	fmt.Println(nums)
	fmt.Println(steps)
}
