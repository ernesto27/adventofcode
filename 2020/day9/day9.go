package day9

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func A() {
	input := []int{}

	file, err := os.Open("day9/input.txt")
	if err != nil {
		fmt.Println(err, "\nYou need a file called input.txt in this directory")
		return
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		number, _ := strconv.Atoi(sc.Text())
		input = append(input, number)
	}

	offset := 25
	preamble := 25
	start := 0
	invalid := 0

out:
	for i := offset; i < len(input); i++ {
		needle := input[i]
		fmt.Println(needle)

		// check sum
		isValid := false
		for x := start; x < preamble; x++ {
			for z := x + 1; z < preamble; z++ {
				if input[x]+input[z] == needle {
					isValid = true
				}
			}
		}

		if !isValid {
			invalid = needle
			break out
		}
		start += 1
		preamble += 1
	}

	fmt.Println("INVALID ", invalid)

	numbersSum := []int{}
out2:
	for x := 0; x < len(input); x++ {
		count := input[x]
		numbersSum = append(numbersSum, input[x])

		for y := x + 1; y < len(input); y++ {
			count += input[y]
			numbersSum = append(numbersSum, input[y])
			if count == invalid {
				break out2
			}

			if count > invalid {
				count = 0
				numbersSum = []int{}
				break
			}
		}
		numbersSum = []int{}
	}
	fmt.Println("NUMBER SUM", numbersSum)

	smallest := math.MaxInt
	largerst := 0
	for _, n := range numbersSum {
		if n < smallest {
			smallest = n
		}
		if n > largerst {
			largerst = n
		}
	}

	fmt.Println(smallest + largerst)

}
