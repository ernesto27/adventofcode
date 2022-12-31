package day08

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func First() {
	items := [][]string{}

	file, err := os.Open("day08/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		currText := sc.Text()
		data := strings.Split(currText, " ")
		item := []string{}
		item = append(item, data[0])
		item = append(item, data[1])
		item = append(item, data[2])
		item = append(item, data[3])
		item = append(item, data[4])
		item = append(item, data[5])
		item = append(item, data[6])
		items = append(items, item)
	}

	registers := make(map[string]int)

	for _, i := range items {
		register := i[0]
		action := i[1]
		registerToCompare := i[4]
		condition := i[5]

		value, err := strconv.Atoi(i[2])
		if err != nil {
			panic(err)
		}

		valueToCompare, err := strconv.Atoi(i[6])
		if err != nil {
			panic(err)
		}

		valueRegisterCompare := registers[registerToCompare]
		isValidCondition := false

		switch condition {
		case ">":
			if valueRegisterCompare > valueToCompare {
				isValidCondition = true
			}
		case "<":
			if valueRegisterCompare < valueToCompare {
				isValidCondition = true
			}
		case ">=":
			if valueRegisterCompare >= valueToCompare {
				isValidCondition = true
			}
		case "<=":
			if valueRegisterCompare <= valueToCompare {
				isValidCondition = true
			}
		case "!=":
			if valueRegisterCompare != valueToCompare {
				isValidCondition = true
			}
		case "==":
			if valueRegisterCompare == valueToCompare {
				isValidCondition = true
			}
		}

		if isValidCondition {
			if action == "inc" {
				registers[register] += value
			} else if action == "dec" {
				registers[register] -= value
			}

			fmt.Println(registers[register])
		}
	}

	fmt.Println(registers)

	max := 0
	for _, v := range registers {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
}

func Second() {
	items := [][]string{}

	file, err := os.Open("day08/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		currText := sc.Text()
		data := strings.Split(currText, " ")
		item := []string{}
		item = append(item, data[0])
		item = append(item, data[1])
		item = append(item, data[2])
		item = append(item, data[3])
		item = append(item, data[4])
		item = append(item, data[5])
		item = append(item, data[6])
		items = append(items, item)
	}

	registers := make(map[string]int)
	maxValue := 0

	for _, i := range items {
		register := i[0]
		action := i[1]
		registerToCompare := i[4]
		condition := i[5]

		value, err := strconv.Atoi(i[2])
		if err != nil {
			panic(err)
		}

		valueToCompare, err := strconv.Atoi(i[6])
		if err != nil {
			panic(err)
		}

		valueRegisterCompare := registers[registerToCompare]
		isValidCondition := false

		switch condition {
		case ">":
			if valueRegisterCompare > valueToCompare {
				isValidCondition = true
			}
		case "<":
			if valueRegisterCompare < valueToCompare {
				isValidCondition = true
			}
		case ">=":
			if valueRegisterCompare >= valueToCompare {
				isValidCondition = true
			}
		case "<=":
			if valueRegisterCompare <= valueToCompare {
				isValidCondition = true
			}
		case "!=":
			if valueRegisterCompare != valueToCompare {
				isValidCondition = true
			}
		case "==":
			if valueRegisterCompare == valueToCompare {
				isValidCondition = true
			}
		}

		if isValidCondition {
			if action == "inc" {
				registers[register] += value
			} else if action == "dec" {
				registers[register] -= value
			}

			if registers[register] > maxValue {
				maxValue = registers[register]
			}
		}
	}

	fmt.Println(maxValue)

}
