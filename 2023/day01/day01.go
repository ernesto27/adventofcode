package day01

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func First() {
	file, err := os.Open("day01/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	var sum int
	for sc.Scan() {
		var pair string
		var last string
		isFirstSet := false
		for _, c := range sc.Text() {
			if unicode.IsDigit(c) {
				ns := fmt.Sprintf("%c", c)
				if !isFirstSet {
					isFirstSet = true
					pair += ns
				}

				last = ns
			}
		}

		pair += last

		n, err := strconv.Atoi(pair)
		if err != nil {
			panic(err)
		}
		sum += n

	}

	fmt.Println("SUM ", sum)
}

type s struct {
	Name     string
	value    string
	Position [][]int
}

type byPosition []s

func (bp byPosition) Len() int      { return len(bp) }
func (bp byPosition) Swap(i, j int) { bp[i], bp[j] = bp[j], bp[i] }
func (bp byPosition) Less(i, j int) bool {
	// Assuming Position is a slice of slices where the first element of each inner slice represents the position
	return bp[i].Position[0][0] < bp[j].Position[0][0]
}

func Second() {
	//wordsToSearch := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	wordsNumber := make(map[string]string)
	wordsNumber["one"] = "1"
	wordsNumber["two"] = "2"
	wordsNumber["three"] = "3"
	wordsNumber["four"] = "4"
	wordsNumber["five"] = "5"
	wordsNumber["six"] = "6"
	wordsNumber["seven"] = "7"
	wordsNumber["eight"] = "8"
	wordsNumber["nine"] = "9"

	file, err := os.Open("day01/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	var sum int
	for sc.Scan() {
		t := sc.Text()

		ss := []s{}

		for k, w := range wordsNumber {
			wordRegex := regexp.MustCompile(k)

			// Find all occurrences of the word in the text
			matches := wordRegex.FindAllStringIndex(t, -1)

			if len(matches) > 0 {
				//resp = append(resp, matches...)
				ss = append(ss, s{
					Name:     k,
					Position: matches,
					value:    w,
				})

			}
		}
		sort.Sort(byPosition(ss))
		for _, st := range ss {
			t = strings.Replace(t, st.Name, st.value, 1)

			//fmt.Println(i, st)
		}

		var pair string
		var last string
		isFirstSet := false
		fmt.Println("line ", t)
		for _, c := range t {
			if unicode.IsDigit(c) {
				ns := fmt.Sprintf("%c", c)
				if !isFirstSet {
					isFirstSet = true
					pair += ns
				}

				last = ns
			}
		}

		pair += last

		fmt.Println("PAIR ", pair)

		n, err := strconv.Atoi(pair)
		if err != nil {
			panic(err)
		}
		sum += n
	}

	fmt.Println("SUM SECOND ", sum)
}
