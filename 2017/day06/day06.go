package day06

import (
	"fmt"
	"strconv"
)

func First() {
	items := []int{2, 8, 8, 5, 4, 2, 3, 1, 5, 5, 1, 2, 15, 13, 5, 14}

	resp := 0
	m := make(map[string]bool)

	for {
		maxValue := items[0]
		maxIndex := 0

		for index, value := range items {
			if value > maxValue {
				maxValue = value
				maxIndex = index
			}
		}

		redistribute := items[maxIndex]
		items[maxIndex] = 0
		for {
			if redistribute == 0 {
				break
			}

			items[(maxIndex+1)%len(items)] += 1
			maxIndex += 1
			redistribute -= 1
		}

		resp += 1

		s := ""
		for _, digit := range items {
			s += strconv.Itoa(digit)
		}

		if _, ok := m[s]; ok {
			break
		}

		m[s] = true

	}

	fmt.Println(resp)
}

func Second() {
	items := []int{2, 8, 8, 5, 4, 2, 3, 1, 5, 5, 1, 2, 15, 13, 5, 14}

	resp := 0
	m := make(map[string]int)
	cycles := 0

	for {
		maxValue := items[0]
		maxIndex := 0

		for index, value := range items {
			if value > maxValue {
				maxValue = value
				maxIndex = index
			}
		}

		redistribute := items[maxIndex]
		items[maxIndex] = 0
		for {
			if redistribute == 0 {
				break
			}

			items[(maxIndex+1)%len(items)] += 1

			maxIndex += 1
			redistribute -= 1
		}

		resp += 1

		s := ""
		for _, digit := range items {
			s += strconv.Itoa(digit)
		}

		if _, ok := m[s]; ok {
			fmt.Println(resp, m[s])
			resp -= m[s]
			break
		}

		cycles += 1
		m[s] = cycles

	}

	fmt.Println(resp)
}
