package day04

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func First() {
	file, _ := os.Open("day04/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	result := 0
	for sc.Scan() {
		currText := sc.Text()
		assigments := strings.Split(currText, ",")

		assigments1 := strings.Split(assigments[0], "-")
		assigments2 := strings.Split(assigments[1], "-")

		pair1 := []int{}
		pair2 := []int{}

		first1, err := strconv.Atoi(assigments1[0])
		if err != nil {
			panic(err)
		}
		first2, err := strconv.Atoi(assigments1[1])
		if err != nil {
			panic(err)
		}

		second1, err := strconv.Atoi(assigments2[0])
		if err != nil {
			panic(err)
		}

		second2, err := strconv.Atoi(assigments2[1])
		if err != nil {
			panic(err)
		}

		if first1 < second1 {
			pair1 = append(pair1, first1)
			pair1 = append(pair1, first2)
			pair2 = append(pair2, second1)
			pair2 = append(pair2, second2)

		} else if first1 == second1 {
			if first2 >= second2 {
				pair1 = append(pair1, first1)
				pair1 = append(pair1, first2)
				pair2 = append(pair2, second1)
				pair2 = append(pair2, second2)
			} else {
				pair1 = append(pair1, second1)
				pair1 = append(pair1, second2)
				pair2 = append(pair2, first1)
				pair2 = append(pair2, first2)
			}

		} else {
			pair1 = append(pair1, second1)
			pair1 = append(pair1, second2)
			pair2 = append(pair2, first1)
			pair2 = append(pair2, first2)
		}

		if pair2[0] >= pair1[0] && pair2[1] <= pair1[1] {
			fmt.Println("found")
			result += 1
		}

		fmt.Println(pair1, pair2)
		fmt.Println("----")
	}

	fmt.Println(result)
}

func Two() {
	file, _ := os.Open("day04/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	result := 0
	for sc.Scan() {
		currText := sc.Text()
		assigments := strings.Split(currText, ",")

		assigments1 := strings.Split(assigments[0], "-")
		assigments2 := strings.Split(assigments[1], "-")

		pair1 := []int{}
		pair2 := []int{}

		first1, err := strconv.Atoi(assigments1[0])
		if err != nil {
			panic(err)
		}
		first2, err := strconv.Atoi(assigments1[1])
		if err != nil {
			panic(err)
		}

		second1, err := strconv.Atoi(assigments2[0])
		if err != nil {
			panic(err)
		}

		second2, err := strconv.Atoi(assigments2[1])
		if err != nil {
			panic(err)
		}

		if first1 < second1 {
			pair1 = append(pair1, first1)
			pair1 = append(pair1, first2)
			pair2 = append(pair2, second1)
			pair2 = append(pair2, second2)

		} else if first1 == second1 {
			if first2 >= second2 {
				pair1 = append(pair1, first1)
				pair1 = append(pair1, first2)
				pair2 = append(pair2, second1)
				pair2 = append(pair2, second2)
			} else {
				pair1 = append(pair1, second1)
				pair1 = append(pair1, second2)
				pair2 = append(pair2, first1)
				pair2 = append(pair2, first2)
			}

		} else {
			pair1 = append(pair1, second1)
			pair1 = append(pair1, second2)
			pair2 = append(pair2, first1)
			pair2 = append(pair2, first2)
		}

		if pair2[0] >= pair1[0] && pair2[0] <= pair1[1] ||
			pair2[1] >= pair1[0] && pair2[1] <= pair1[1] {
			fmt.Println("found")
			result += 1
		}

		fmt.Println(pair1, pair2)
		fmt.Println("----")
	}

	fmt.Println(result)
}
