package day05

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func First() {
	file, _ := os.Open("day05/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	crates := [][]string{}
	data := []string{"Z", "J", "G"}
	crates = append(crates, data)
	data = []string{"Q", "L", "R", "P", "W", "F", "V", "C"}
	crates = append(crates, data)
	data = []string{"F", "P", "M", "C", "L", "G", "R"}
	crates = append(crates, data)
	data = []string{"L", "F", "B", "W", "P", "H", "M"}
	crates = append(crates, data)
	data = []string{"G", "C", "F", "S", "V", "Q"}
	crates = append(crates, data)
	data = []string{"W", "H", "J", "Z", "M", "Q", "T", "L"}
	crates = append(crates, data)
	data = []string{"H", "F", "S", "B", "V"}
	crates = append(crates, data)
	data = []string{"F", "J", "Z", "S"}
	crates = append(crates, data)
	data = []string{"M", "C", "D", "P", "F", "H", "B", "T"}
	crates = append(crates, data)

	fmt.Println("CRATES ", crates)

	re := regexp.MustCompile(`\d+`)
	lineCount := 0

	for sc.Scan() {
		lineCount++

		if lineCount >= 11 {
			currText := sc.Text()

			numbers := re.FindAllString(currText, -1)

			arragment := []int{}
			for _, number := range numbers {
				n, err := strconv.Atoi(number)
				if err != nil {
					panic(err)
				}
				arragment = append(arragment, n)
			}

			fmt.Println(currText)
			fmt.Println(arragment)

			for x := 1; x <= arragment[0]; x++ {
				fmt.Println(x)

				// removes item from
				l := len(crates[arragment[1]-1])
				latestFrom := crates[arragment[1]-1][l-1]
				crates[arragment[1]-1] = crates[arragment[1]-1][:l-1]

				// add item to
				crates[arragment[2]-1] = append(crates[arragment[2]-1], latestFrom)
			}

			fmt.Println("------")
		}
	}

	fmt.Println("CRATES LAST ", crates)

	resp := ""
	for _, c := range crates {
		resp += c[len(c)-1]
	}
	fmt.Println(resp)
}

func Two() {
	file, _ := os.Open("day05/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	crates := [][]string{}
	data := []string{"Z", "J", "G"}
	crates = append(crates, data)
	data = []string{"Q", "L", "R", "P", "W", "F", "V", "C"}
	crates = append(crates, data)
	data = []string{"F", "P", "M", "C", "L", "G", "R"}
	crates = append(crates, data)
	data = []string{"L", "F", "B", "W", "P", "H", "M"}
	crates = append(crates, data)
	data = []string{"G", "C", "F", "S", "V", "Q"}
	crates = append(crates, data)
	data = []string{"W", "H", "J", "Z", "M", "Q", "T", "L"}
	crates = append(crates, data)
	data = []string{"H", "F", "S", "B", "V"}
	crates = append(crates, data)
	data = []string{"F", "J", "Z", "S"}
	crates = append(crates, data)
	data = []string{"M", "C", "D", "P", "F", "H", "B", "T"}
	crates = append(crates, data)

	re := regexp.MustCompile(`\d+`)
	lineCount := 0

	for sc.Scan() {
		lineCount++
		// fmt.Println(lineCount)

		if lineCount >= 11 {
			currText := sc.Text()

			numbers := re.FindAllString(currText, -1)

			arragment := []int{}
			for _, number := range numbers {
				n, err := strconv.Atoi(number)
				if err != nil {
					panic(err)
				}
				arragment = append(arragment, n)
			}

			fmt.Println(currText)
			fmt.Println(arragment)

			itemsToMove := []string{}
			for x := 1; x <= arragment[0]; x++ {
				fmt.Println(x)

				// removes item from
				l := len(crates[arragment[1]-1])
				latestFrom := crates[arragment[1]-1][l-1]
				crates[arragment[1]-1] = crates[arragment[1]-1][:l-1]

				// add item to
				itemsToMove = append(itemsToMove, latestFrom)
				// crates[arragment[2]-1] = append(crates[arragment[2]-1], latestFrom)
			}

			for i := len(itemsToMove) - 1; i >= 0; i-- {
				crates[arragment[2]-1] = append(crates[arragment[2]-1], itemsToMove[i])
			}
			fmt.Println("------")
		}
	}

	fmt.Println("CRATES LAST ", crates)

	resp := ""
	for _, c := range crates {
		resp += c[len(c)-1]
	}
	fmt.Println(resp)
}
