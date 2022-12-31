package day12

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func First() {
	items := make(map[int][]int)
	file, err := os.Open("day12/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		getItems(sc.Text(), &items)
	}

	graph := items

	result := bfs(graph, 0)
	fmt.Println(len(result))

}

func Second() {
	items := make(map[int][]int)
	file, err := os.Open("day12/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		getItems(sc.Text(), &items)
	}

	graph := items
	found := make(map[int]bool)
	response := 0

	for x := 0; x < len(items); x++ {
		result := make(map[int]bool)
		if _, ok := found[x]; !ok {
			result = bfs(graph, x)
			response += 1
		}

		for k, _ := range result {
			found[k] = true
		}
	}
	fmt.Println(response)
}

func getItems(text string, items *map[int][]int) {
	parts := strings.Split(text, "<->")

	first := parts[0]
	first = strings.TrimSpace(first)

	num, err := strconv.Atoi(first)
	if err != nil {
		return
	}

	parts = strings.Split(parts[1], ",")

	numbers := []int{}
	for _, part := range parts {
		part = strings.TrimSpace(part)
		numChild, err := strconv.Atoi(part)
		if err != nil {
			continue
		}
		numbers = append(numbers, numChild)
	}

	(*items)[num] = numbers

}

func bfs(graph map[int][]int, start int) map[int]bool {
	queue := []int{}

	visited := map[int]bool{}

	queue = append(queue, start)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbor := range graph[current] {
			if _, ok := visited[neighbor]; !ok {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}
	return visited
}
