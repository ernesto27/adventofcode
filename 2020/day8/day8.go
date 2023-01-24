package day8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func A() {
	input := []string{}

	file, err := os.Open("day8/input.txt")
	if err != nil {
		fmt.Println(err, "\nYou need a file called input.txt in this directory")
		return
	}
	defer file.Close()

	// --- START --
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		input = append(input, sc.Text())
	}

	instructions := make(map[string]bool)
	index := 0
	accumulator := 0
	for {
		// check repeat map
		if _, ok := instructions[fmt.Sprint(index)+input[index]]; ok {
			break
		}

		instructions[fmt.Sprint(index)+input[index]] = true
		code := input[index][:3]
		fmt.Println(code)
		argument := strings.Split(input[index], " ")

		sign := string(argument[1][0])
		fmt.Println(sign)
		number := argument[1][1:]
		n, _ := strconv.Atoi(number)
		fmt.Println(n)

		fmt.Println("---")

		switch code {
		case "acc":
			if sign == "+" {
				accumulator += n
			} else if sign == "-" {
				accumulator -= n
			}
			index++
		case "jmp":
			if sign == "+" {
				index += n
			} else if sign == "-" {
				index -= n
			}
		case "nop":
			index++
		}
	}

	fmt.Println(accumulator)
}

func B() {

	input := []string{}

	file, err := os.Open("day8/input.txt")
	if err != nil {
		fmt.Println(err, "\nYou need a file called input.txt in this directory")
		return
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		input = append(input, sc.Text())
	}

	c := make([]string, len(input))
	copy(c, input)

	for z, i := range input {
		fmt.Println(i)
		code := i[:3]

		c := make([]string, len(input))
		copy(c, input)

		switch code {
		case "nop":
			c[z] = "jmp" + c[z][3:]
		case "jmp":
			c[z] = "nop" + c[z][3:]

		}
		fmt.Println("check if operation terminates")
		fmt.Println(c)

		instructions := make(map[string]bool)
		index := 0
		accumulator := 0
		finish := false
		for {

			if index == len(c) {
				finish = true
				break
			}

			// check repeat map
			if _, ok := instructions[fmt.Sprint(index)+c[index]]; ok {
				break
			}

			instructions[fmt.Sprint(index)+c[index]] = true
			code := c[index][:3]
			fmt.Println(code)
			argument := strings.Split(c[index], " ")

			sign := string(argument[1][0])
			fmt.Println(sign)
			number := argument[1][1:]
			n, _ := strconv.Atoi(number)
			fmt.Println(n)

			fmt.Println("---")

			switch code {
			case "acc":
				if sign == "+" {
					accumulator += n
				} else if sign == "-" {
					accumulator -= n
				}
				index++
			case "jmp":
				if sign == "+" {
					index += n
				} else if sign == "-" {
					index -= n
				}
			case "nop":
				index++
			}
		}

		if finish {
			fmt.Println("ACUMULATOR ", accumulator)
			break
		}

		fmt.Println("----")

	}

}
