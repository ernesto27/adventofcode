package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func First() {
	file, err := os.Open("day03/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var res int
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			fmt.Println(match)
			if len(match) == 3 {
				num1, err := strconv.Atoi(match[1])
				if err != nil {
					panic(err)
				}
				num2, err := strconv.Atoi(match[2])
				if err != nil {
					panic(err)
				}
				res += num1 * num2
			}
		}
	}

	fmt.Println(res)
}

func Second() {
	file, err := os.Open("day03/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var res int
	reDo := regexp.MustCompile(`do\(\)`)
	reMul := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	for scanner.Scan() {
		line := scanner.Text()
		mulMatches := reMul.FindAllStringSubmatchIndex(line, -1)
		doMatches := reDo.FindAllStringSubmatchIndex(line, -1)
		matches := reMul.FindAllStringSubmatch(line, -1)
		matchesDo := reDo.FindAllStringSubmatch(line, -1)

		// fmt.Println(doMatches)
		fmt.Println(mulMatches)
		fmt.Println(doMatches)
		fmt.Println(matches)
		fmt.Println(matchesDo)

		// for _, mulMatch := range mulMatches {
		// 	mulStart := mulMatch[0]
		// 	valid := false

		// 	for _, doMatch := range doMatches {
		// 		doEnd := doMatch[1]
		// 		if doEnd < mulStart {
		// 			valid = true
		// 			break
		// 		}
		// 	}

		// 	if valid {
		// 		num1, err1 := strconv.Atoi(line[mulMatch[2]:mulMatch[3]])
		// 		num2, err2 := strconv.Atoi(line[mulMatch[4]:mulMatch[5]])
		// 		if err1 != nil || err2 != nil {
		// 			fmt.Println("Error converting string to int:", err1, err2)
		// 			continue
		// 		}
		// 		res += num1 * num2
		// 	}
		// }
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(res)
}
