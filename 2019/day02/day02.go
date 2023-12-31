package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func First() {
	data, err := os.ReadFile("day02/input.txt")
	if err != nil {
		panic(err)
	}

	d := strings.Split(string(data), ",")

	for i := 0; i < len(d); i += 4 {
		opCode := d[i]
		a := d[i+1]
		b := d[i+2]

		first, err := strconv.Atoi(a)
		must(err)
		second, err := strconv.Atoi(b)
		must(err)

		val1, err := strconv.Atoi(d[first])
		must(err)
		val2, err := strconv.Atoi(d[second])
		must(err)

		if opCode == "1" {

			sum := val1 + val2
			dest, err := strconv.Atoi(d[i+3])
			must(err)
			d[dest] = fmt.Sprint(sum)
		} else if opCode == "2" {
			res := val1 * val2
			dest, err := strconv.Atoi(d[i+3])
			must(err)

			d[dest] = fmt.Sprint(res)
		} else if opCode == "99" {
			break
		}
	}

	fmt.Println(d)
}
