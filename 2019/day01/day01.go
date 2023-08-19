package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func First() {
	file, err := os.Open("day01/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	sc := bufio.NewScanner(file)
	total := 0

	for sc.Scan() {
		currText := sc.Text()

		n, err := strconv.Atoi(currText)
		if err != nil {
			panic(err)
		}

		mass := n/3 - 2
		total += mass
	}

	fmt.Println(total)
	fmt.Println(100756 / 2)
}

func Second() {

	file, err := os.Open("day01/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	sc := bufio.NewScanner(file)
	total := 0

	for sc.Scan() {
		currText := sc.Text()

		n, err := strconv.Atoi(currText)
		if err != nil {
			panic(err)
		}

		for {
			mass := n/3 - 2
			if mass <= 0 {
				break
			}
			n = mass
			total += mass
		}
	}
	fmt.Println(total)

}
