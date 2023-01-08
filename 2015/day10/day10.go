package day10

import (
	"fmt"
	"strconv"
)

func recursive(input string, step int) {
	if step == 0 {
		fmt.Println(len(input))
		return
	}

	var latest rune
	latest = -99999
	count := 1
	result := ""
	for index, i := range input {
		if i == latest {
			count++
		} else {

			if index != 0 {
				result += fmt.Sprintf("%d%c", count, latest)
			}
			count = 1
			latest = i

		}

		if index == len(input)-1 {
			result += fmt.Sprintf("%d%c", count, latest)

		}
	}

	input = result
	recursive(input, step-1)

}

func First() {
	input := "1113222113"
	recursive(input, 50)

}

func Second() {

	var numbers []byte = []byte("1113222113")

	for i := 0; i < 50; i++ {
		var newNumbers []byte

		for j := 0; j < len(numbers); j++ {
			num := numbers[j]

			start := j
			for j < len(numbers)-1 && numbers[j+1] == num {
				j++
			}
			end := j

			length := end - start + 1

			newNumbers = append(newNumbers, strconv.Itoa(length)[0])
			newNumbers = append(newNumbers, numbers[j])
		}

		numbers = newNumbers
	}

	println(len(numbers))
}
