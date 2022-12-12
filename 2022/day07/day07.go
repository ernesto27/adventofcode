package day07

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Tree struct {
	Root *Node
}

type Node struct {
	Name     string
	Value    int
	Children []*Node
	Prev     *Node
}

func traverse(node *Node, result *int) {
	fmt.Println("CHECK", node.Value)
	if node.Value <= 100000 {
		*result += node.Value
	}

	for _, child := range node.Children {
		traverse(child, result)
	}
}

func traverseTwo(node *Node, smallestRes *int, minRequired int) {
	if node.Value >= minRequired && node.Value < *smallestRes {
		*smallestRes = node.Value
	}

	for _, child := range node.Children {
		traverseTwo(child, smallestRes, minRequired)
	}
}

func First() {

	file, _ := os.Open("day07/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	tree := &Tree{}
	root := &Node{Name: "root", Value: 0}
	tree.Root = root
	current := root

	for sc.Scan() {
		currText := sc.Text()
		items := strings.Split(currText, " ")

		if items[0] == "$" {
			// command
			if items[1] == "cd" {
				switch items[2] {
				case "/":
					current = root
				case "..":
					// go back
					current = current.Prev
				default:
					// go to new dir
					fmt.Println(items)
					n := &Node{Name: items[2], Prev: current}
					current.Children = append(current.Children, n)
					current = n
				}
			}
		} else {
			number, err := strconv.Atoi(items[0])
			if err == nil {
				// number add sum
				current.Value += number
				if current.Prev != nil {
					prev := current.Prev
					for prev != nil {
						prev.Value += number
						prev = prev.Prev
					}
				}
			}
		}
	}

	result := 0
	traverse(tree.Root, &result)
	fmt.Println(result)
}

func Two() {

	file, _ := os.Open("day07/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	tree := &Tree{}
	root := &Node{Name: "root", Value: 0}
	tree.Root = root
	current := root

	for sc.Scan() {
		currText := sc.Text()
		items := strings.Split(currText, " ")

		if items[0] == "$" {
			// command
			if items[1] == "cd" {
				switch items[2] {
				case "/":
					current = root
				case "..":
					// go back
					current = current.Prev
				default:
					// go to new dir
					n := &Node{Name: items[2], Prev: current}
					current.Children = append(current.Children, n)
					current = n
				}
			}
		} else {
			number, err := strconv.Atoi(items[0])
			if err == nil {
				// number add sum
				current.Value += number
				if current.Prev != nil {
					prev := current.Prev
					for prev != nil {
						prev.Value += number
						prev = prev.Prev
					}
				}
			}
		}
	}

	totalSpace := 70000000
	unusedSpace := totalSpace - root.Value
	spaceRequired := 30000000
	minRequired := spaceRequired - unusedSpace
	smallestRes := math.MaxInt
	traverseTwo(tree.Root, &smallestRes, minRequired)
	fmt.Println(smallestRes)
}
