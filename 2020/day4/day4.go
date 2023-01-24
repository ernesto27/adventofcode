package day4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func First() {
	required := make(map[string]bool)
	required["byr"] = true
	required["iyr"] = true
	required["eyr"] = true
	required["hgt"] = true
	required["hcl"] = true
	required["ecl"] = true
	required["pid"] = true

	file, _ := os.Open("day4/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	var input [][]string
	temp := []string{}
	for sc.Scan() {
		t := sc.Text()

		if t == "" {
			input = append(input, temp)
			temp = []string{}
			fmt.Println("----")
		} else {
			s := strings.Split(t, " ")
			for _, v := range s {
				fmt.Println(v)
				temp = append(temp, v)
			}
		}
	}
	input = append(input, temp)

	fmt.Println(input[0])

	valid := 0
	for _, i := range input {
		count := 0
		for _, j := range i {
			fmt.Println(j)
			s := strings.Split(j, ":")
			// check map key
			if _, ok := required[s[0]]; ok {
				count++
			}
		}

		if count == 7 {
			valid++
		}
		fmt.Println("----")
	}

	fmt.Println(valid)
}

func Second() {
	required := make(map[string]func(value string) bool)
	required["byr"] = func(value string) bool {
		i, err := strconv.Atoi(value)
		if err != nil {
			return false
		}

		if i >= 1920 && i <= 2002 {
			return true
		}

		return false
	}
	required["iyr"] = func(value string) bool {
		i, err := strconv.Atoi(value)
		if err != nil {
			return false
		}

		if i >= 2010 && i <= 2020 {
			return true
		}

		return false
	}
	required["eyr"] = func(value string) bool {
		i, err := strconv.Atoi(value)
		if err != nil {
			return false
		}

		if i >= 2020 && i <= 2030 {
			return true
		}

		return false
	}
	required["hgt"] = func(value string) bool {
		i, err := strconv.Atoi(value[:len(value)-2])
		typeH := value[len(value)-2:]
		if err != nil {
			return false
		}

		if typeH == "cm" {
			if i >= 150 && i <= 193 {
				return true
			}
		} else if typeH == "in" {
			if i >= 59 && i <= 76 {
				return true
			}
		}

		return false
	}
	required["hcl"] = func(value string) bool {
		if value[0] == '#' {
			chars := value[1:]
			fmt.Println(chars)
			// regexp [0-9a-f]{6}
			m, err := regexp.MatchString(`[0-9a-f]{6}`, chars)
			if err != nil {
				return false
			}
			if m {
				return true
			}
		}

		return false
	}
	required["ecl"] = func(value string) bool {
		// check string in array of string
		ecl := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, e := range ecl {
			if e == value {
				return true
			}
		}
		return false
	}
	required["pid"] = func(value string) bool {
		_, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		if len(value) == 9 {
			return true
		}
		return false
	}

	file, _ := os.Open("day4/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	var input [][]string
	temp := []string{}
	for sc.Scan() {
		t := sc.Text()

		if t == "" {
			input = append(input, temp)
			temp = []string{}
			fmt.Println("----")
		} else {
			s := strings.Split(t, " ")
			for _, v := range s {
				fmt.Println(v)
				temp = append(temp, v)
			}
		}
	}
	input = append(input, temp)

	fmt.Println(input[0])

	valid := 0
	for _, i := range input {
		count := 0
		for _, j := range i {
			fmt.Println(j)
			s := strings.Split(j, ":")
			// check map key
			if _, ok := required[s[0]]; ok {
				f := required[s[0]]
				if f(s[1]) {
					count++
				}
			}
		}

		if count == 7 {
			valid++
		}
		fmt.Println("----")
	}

	fmt.Println(valid)
}
