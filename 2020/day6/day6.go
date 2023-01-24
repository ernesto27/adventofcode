package day6

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func First() {
	var input [][][]string

	file, err := os.Open("day6/input.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("You need a file called input.txt in this directory")
		return
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	temp := [][]string{}
	for sc.Scan() {
		if sc.Text() == "" {
			fmt.Println("----")
			input = append(input, temp)
			temp = [][]string{}
		}

		s := strings.Split(sc.Text(), "")
		x := []string{}
		for _, c := range s {
			x = append(x, c)
		}
		temp = append(temp, x)
	}
	input = append(input, temp)

	fmt.Println(input[0])
	fmt.Println(input[len(input)-1])

	sum := 0
	for _, i := range input {
		chars := make(map[string]bool)
		for _, j := range i {
			for _, k := range j {
				_, ok := chars[k]
				if !ok {
					chars[k] = true

				}
				fmt.Print(k)
			}
			fmt.Println()
		}
		fmt.Println("----")
		sum += len(chars)
	}

	fmt.Println(sum)
}

func Second() {
	var input [][][]string

	file, err := os.Open("day6/input.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("You need a file called input.txt in this directory")
		return
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	temp := [][]string{}
	for sc.Scan() {
		if sc.Text() == "" {
			fmt.Println("----")
			input = append(input, temp)
			// empty slice
			temp = [][]string{}
			continue
		}

		s := strings.Split(sc.Text(), "")
		x := []string{}
		for _, c := range s {
			x = append(x, c)
		}
		temp = append(temp, x)
	}
	input = append(input, temp)

	sum := 0
	for _, i := range input {
		chars := make(map[string]int)
		persons := 0
		for _, j := range i {
			persons++
			for _, k := range j {
				_, ok := chars[k]
				if ok {
					chars[k]++
				} else {
					chars[k] = 1
				}
				fmt.Print(k)
			}
			fmt.Println()
		}
		fmt.Println(persons)
		fmt.Println("----")
		for _, v := range chars {
			if v == persons {
				sum++
			}
		}
	}

	fmt.Println("SUM ", sum)
}
