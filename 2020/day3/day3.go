package day3

import (
	"bufio"
	"fmt"
	"os"
)

func First() {

	var input []string

	file, _ := os.Open("day3/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		input = append(input, sc.Text())
	}

	rigthMove := 3
	count := 0
	for index, val := range input {
		if index != 0 {
			if rigthMove > len(val)-1 {
				fmt.Println("change right move")
				rigthMove = rigthMove - len(val)
			}

			if string(val[rigthMove]) == "#" {
				fmt.Println("X")
				count++
			} else {
				fmt.Println("0")
			}
			rigthMove += 3
		}
	}

	fmt.Println(count)
}

func Second() {
	var input []string
	file, _ := os.Open("day3/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		input = append(input, sc.Text())
	}

	moves := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	counts := []int{}
	for _, m := range moves {
		rigthMove := m[0]
		downMove := m[1]
		count := 0

		for x := 0; x < len(input); x += downMove {
			if x != 0 {
				if rigthMove > len(input[x])-1 {
					rigthMove = rigthMove - len(input[x])
				}

				if string(input[x][rigthMove]) == "#" {
					count++
				} else {
					fmt.Println("0")
				}
				rigthMove += m[0]
			}
		}
		counts = append(counts, count)
	}

	fmt.Println(counts)
	result := 0
	for i := 0; i < len(counts); i++ {
		if i == 0 {
			result = counts[i] * counts[i+1]
			i += 1
		} else {
			result = result * counts[i]
		}
	}

	fmt.Println(result)
}
