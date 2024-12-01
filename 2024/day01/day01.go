package day01

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func First() {
	// left := []int{3, 4, 2, 1, 3, 3}
	// right := []int{4, 3, 5, 3, 9, 3}

	file, err := os.Open("day01/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	left := []int{}
	right := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		if len(numbers) == 2 {
			leftNum, err1 := strconv.Atoi(numbers[0])
			rightNum, err2 := strconv.Atoi(numbers[1])
			if err1 == nil && err2 == nil {
				left = append(left, leftNum)
				right = append(right, rightNum)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	var count int
	fmt.Println(left, right)
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff < 0 {
			diff = -diff
		}
		count += diff
	}

	fmt.Println(count)
}

func Second() {
	// left := []int{3, 4, 2, 1, 3, 3}
	// right := []int{4, 3, 5, 3, 9, 3}

	file, err := os.Open("day01/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	left := []int{}
	right := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		if len(numbers) == 2 {
			leftNum, err1 := strconv.Atoi(numbers[0])
			rightNum, err2 := strconv.Atoi(numbers[1])
			if err1 == nil && err2 == nil {
				left = append(left, leftNum)
				right = append(right, rightNum)
			}
		}
	}

	type data struct {
		leftCount  int
		rightCount int
	}
	locations := make(map[int]data)

	// for _, l := range left {
	// 	var currCount int
	// 	if _, ok := locations[l]; ok {
	// 		currCount = locations[l].leftCount + 1
	// 	} else {
	// 		currCount = 1
	// 	}

	// 	locations[l] = data{
	// 		leftCount: currCount,
	// 	}
	// }

	// for _, r := range right {
	// 	var currCount int
	// 	if _, ok := locations[r]; ok {
	// 		currCount = locations[r].rightCount + 1
	// 	} else {
	// 		continue
	// 	}

	// 	data := locations[r]
	// 	data.rightCount = currCount
	// 	locations[r] = data
	// }

	for i := 0; i < len(left); i++ {
		l := left[i]
		r := right[i]
		if loc, exists := locations[l]; exists {
			loc.leftCount++
			locations[l] = loc
		} else {
			locations[l] = data{leftCount: 1}
		}
		if loc, exists := locations[r]; exists {
			loc.rightCount++
			locations[r] = loc
		} else {
			locations[r] = data{rightCount: 1}
		}
	}

	var result int

	for k, v := range locations {
		for i := 0; i < v.leftCount; i++ {
			result += k * v.rightCount
		}
	}

	fmt.Println(locations)
	fmt.Println(result)
}
