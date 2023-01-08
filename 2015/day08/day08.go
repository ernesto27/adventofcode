package day08

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func First() {
	input, err := ioutil.ReadFile("day08/raw.txt")
	if err != nil {
		panic(err)
	}

	i := strings.Split(string(input), "\n")

	codeChars := 0
	memChars := 0
	for _, item := range i {
		fmt.Println(item)
		fmt.Println(len(item))
		codeChars += len(item)
		unescaped, _ := strconv.Unquote(item)
		memChars += len(unescaped)
	}

	fmt.Println(codeChars - memChars)
}

func Second() {
	input, err := ioutil.ReadFile("day08/raw.txt")
	if err != nil {
		panic(err)
	}

	i := strings.Split(string(input), "\n")

	codeChars := 0
	encodedChars := 0
	for _, item := range i {
		fmt.Println(item)
		fmt.Println(len(item))
		codeChars += len(item)
		escaped := strconv.Quote(item)
		encodedChars += len(escaped)
	}

	fmt.Println(encodedChars - codeChars)
}
