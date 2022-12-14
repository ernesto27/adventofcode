package day16

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func First() {
	input, err := ioutil.ReadFile("day16/day16.txt")
	if err != nil {
		panic(err)
	}

	sues := strings.Split(string(input), "\n")

	compounds := map[string]string{
		"children":    "3",
		"cats":        "7",
		"samoyeds":    "2",
		"pomeranians": "3",
		"akitas":      "0",
		"vizslas":     "0",
		"goldfish":    "5",
		"trees":       "3",
		"cars":        "2",
		"perfumes":    "1",
	}

	regexes := map[string]*regexp.Regexp{
		"number":      regexp.MustCompile("Sue (\\d+)"),
		"children":    regexp.MustCompile("children: (\\d+)"),
		"cats":        regexp.MustCompile("cats: (\\d+)"),
		"samoyeds":    regexp.MustCompile("samoyeds: (\\d+)"),
		"pomeranians": regexp.MustCompile("pomeranians: (\\d+)"),
		"akitas":      regexp.MustCompile("akitas: (\\d+)"),
		"vizslas":     regexp.MustCompile("vizslas: (\\d+)"),
		"goldfish":    regexp.MustCompile("goldfish: (\\d+)"),
		"trees":       regexp.MustCompile("trees: (\\d+)"),
		"cars":        regexp.MustCompile("cars: (\\d+)"),
		"perfumes":    regexp.MustCompile("perfumes: (\\d+)"),
	}

	for _, sue := range sues {
		realSue := true

		for compound, amount := range compounds {
			match := regexes[compound].FindStringSubmatch(sue)
			if len(match) > 0 && match[1] != amount {
				realSue = false
				break
			}
		}

		if realSue {
			println(regexes["number"].FindStringSubmatch(sue)[1])
			break
		}
	}
}

func Second() {
	input, err := ioutil.ReadFile("day16/day16.txt")
	if err != nil {
		panic(err)
	}

	sues := strings.Split(string(input), "\n")

	compounds := map[string]string{
		"children":    "3",
		"cats":        "7",
		"samoyeds":    "2",
		"pomeranians": "3",
		"akitas":      "0",
		"vizslas":     "0",
		"goldfish":    "5",
		"trees":       "3",
		"cars":        "2",
		"perfumes":    "1",
	}

	regexes := map[string]*regexp.Regexp{
		"number":      regexp.MustCompile("Sue (\\d+)"),
		"children":    regexp.MustCompile("children: (\\d+)"),
		"cats":        regexp.MustCompile("cats: (\\d+)"),
		"samoyeds":    regexp.MustCompile("samoyeds: (\\d+)"),
		"pomeranians": regexp.MustCompile("pomeranians: (\\d+)"),
		"akitas":      regexp.MustCompile("akitas: (\\d+)"),
		"vizslas":     regexp.MustCompile("vizslas: (\\d+)"),
		"goldfish":    regexp.MustCompile("goldfish: (\\d+)"),
		"trees":       regexp.MustCompile("trees: (\\d+)"),
		"cars":        regexp.MustCompile("cars: (\\d+)"),
		"perfumes":    regexp.MustCompile("perfumes: (\\d+)"),
	}

	for _, sue := range sues {
		realSue := true

		for compound, amount := range compounds {
			match := regexes[compound].FindStringSubmatch(sue)
			if compound == "cats" || compound == "trees" {
				if len(match) > 0 && match[1] <= amount {
					realSue = false
					break
				}
			} else if compound == "pomeranians" || compound == "goldfish" {
				if len(match) > 0 && match[1] >= amount {
					realSue = false
					break
				}
			} else {
				if len(match) > 0 && match[1] != amount {
					realSue = false
					break
				}
			}

		}

		if realSue {
			println(regexes["number"].FindStringSubmatch(sue)[1])
			break
		}
	}
}
