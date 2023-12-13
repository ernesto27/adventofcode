package main

import "fmt"

func First() {

	i := []string{"L", "L", "R"}
	m := make(map[string][]string)
	m["AAA"] = []string{"BBB", "BBB"}
	m["BBB"] = []string{"AAA", "ZZZ"}
	m["ZZZ"] = []string{"ZZZ", "ZZZ"}

	index := 1
	count := 1
	curr := m["AAA"][0]
	if curr == "ZZZ" {
		fmt.Println(index)
	}

	for {
		if index >= len(i) {
			index = 0
		}

		s := i[index]

		pos := 0
		if s == "R" {
			pos = 1
		}

		cm := m[curr]
		curr = cm[pos]

		index += 1
		count += 1

		if curr == "ZZZ" {
			fmt.Println(count)
			fmt.Println("found")
			break
		}

	}
}

func main() {
	First()
}
