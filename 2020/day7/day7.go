package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func lookIn(lookedBag string, rules *map[string][]string, valids *map[string]bool) bool {
	if (*valids)[lookedBag] {
		return true
	}
	for _, bag := range (*rules)[lookedBag] { //I lookIn each bag contained in the current bag recursively
		if lookIn(bag, rules, valids) {
			(*valids)[bag] = true
			return true
		}
	}
	return false
}

type bagLabel struct {
	bag    string
	number int
}

func countIn(lookedbag string, rules *map[string][]bagLabel) int {
	countBags := 1
	for _, bag := range (*rules)[lookedbag] { //I countIn each bag contained in the current bag recursively
		countBags += bag.number * countIn(bag.bag, rules)
	}
	return countBags
}

func B() {
	input, err := os.Open("day7/input.txt")
	if err != nil {
		fmt.Println(err, "\nYou need a file called input.txt in this directory")
		return
	}
	defer input.Close()

	//	--- START ---
	sc := bufio.NewScanner(input)
	rules := make(map[string][]bagLabel) //each bag is mapped to the list of bags that contains with the number of istances of each bag

	for sc.Scan() {
		parsedInput := strings.ReplaceAll(sc.Text(), " bags", "")
		parsedInput = strings.ReplaceAll(parsedInput, " bag", "")
		parsedInput = strings.ReplaceAll(parsedInput, ".", "")
		rule := strings.Split(parsedInput, " contain ")

		for _, bag := range strings.Split(rule[1], ", ") {
			numberOfBags, _ := strconv.Atoi(bag[0:1])
			rules[rule[0]] = append(rules[rule[0]], bagLabel{bag[2:], numberOfBags})
		}
	}
	fmt.Println(countIn("shiny gold", &rules) - 1) // -1 because I don't have to count the Shiny gold bag
}

func A() {
	input, err := os.Open("day7/input.txt")
	if err != nil {
		fmt.Println(err, "\nYou need a file called input.txt in this directory")
		return
	}
	defer input.Close()

	// --- START --
	sc := bufio.NewScanner(input)
	rules := make(map[string][]string) //each bag is mapped to the list of bags that contains
	valids := make(map[string]bool)    //each bag that contains a shiny gold bag is a valid bag
	valids["shiny gold"] = true

	for sc.Scan() {
		parsedInput := strings.ReplaceAll(sc.Text(), " bags", "")
		parsedInput = strings.ReplaceAll(parsedInput, " bag", "")
		parsedInput = strings.ReplaceAll(parsedInput, ".", "")
		rule := strings.Split(parsedInput, " contain ")

		for _, bag := range strings.Split(rule[1], ", ") {
			rules[rule[0]] = append(rules[rule[0]], bag[2:])
		}
	}

	for bag := range rules {
		if lookIn(bag, &rules, &valids) { //lookIn says if the bag contains a shiny gold bag
			valids[bag] = true
		}
	}
	fmt.Println(len(valids) - 1)
}

type Bag struct {
	Name     string
	Children []*Bag
}

func search(items []map[string]int, bags map[string][]map[string]int) bool {
	fmt.Println("CALL SEARCH")
	found := false

out:
	for _, i := range items {
		for z, _ := range i {
			fmt.Println("SEARCH", z)
			if z == "SHINY_GOLD" {
				fmt.Println("FOUND")
				found = true
				break out
			} else {
				// check key on map
				if !found {
					if _, ok := bags[z]; ok {
						search(bags[z], bags)
					}

				}
			}

		}
	}

	return found
}

func First() {

	bags := make(map[string][]map[string]int)

	bags["LIGTH_RED"] = []map[string]int{
		{"BRIGTH_WHITE": 1},
		{"MUTED_YELLOW": 2},
	}
	bags["BRIGTH_WHITE"] = []map[string]int{
		{"SHINY_GOLD": 1},
	}
	bags["DARK_ORANGE"] = []map[string]int{
		{"BRIGTH_WHITE": 3},
		{"MUTED_YELLOW": 4},
	}

	for key, value := range bags {
		fmt.Println(key, value)
		fmt.Println("-----")
		resp := search(value, bags)
		fmt.Println("FOUND ", resp)
		break
	}

	return

	bagsMap := make(map[string][]map[string]int)
	file, err := os.Open("day7/raw.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("You need a file called input.txt in this directory")
		return
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		// Split by ,  means bags
		bags := strings.Split(sc.Text(), " contain ")
		// fmt.Println(bags[0])
		keyMap := strings.Replace(bags[0], " ", "_", -1)
		// fmt.Println(keyMap)

		// children
		mapChildren := []map[string]int{}
		children := strings.Split(bags[1], ",")
		// fmt.Println(children)
		for _, c := range children {
			if string(c[0]) == " " {
				c = c[1:]
			}
			count := string(c[0])
			n, _ := strconv.Atoi(count)
			// remove spaces
			c = c[2:]
			c = strings.Replace(c, " ", "_", -1)
			c = strings.Replace(c, ".", "", -1)
			// fmt.Println(c)
			mapChildren = append(mapChildren, map[string]int{
				c: n,
			})
		}
		bagsMap[keyMap] = mapChildren
		// fmt.Println("----")
	}

	count := 0
	for key, values := range bagsMap {
		fmt.Println(key)
		// fmt.Println(values)
		for _, v := range values {
			for y, _ := range v {
				fmt.Println(y)
				if y == "shiny_gold_bags" {
					count++
					break
				}
			}
			fmt.Println("-----")
		}
		break
	}

	fmt.Println(count)
	return

	///////////////////////////////////////////////////////////

	n := Bag{}
	n.Name = "LIGTH_RED"
	n.Children = []*Bag{
		{
			Name: "BRIGTH_WHITE",
			Children: []*Bag{
				{
					Name: "SHINY_GOLD",
					Children: []*Bag{
						{
							Name: "NEW ITEM",
						},
					},
				},
			},
		},
		{
			Name: "MUTE_YELLOW",
			Children: []*Bag{
				{
					Name: "SHINY_GOLD",
				},
				{
					Name: "FADED_BLUE",
				},
			},
		},
	}

	// loop recursive on Bags
	// Searc	h(&n)
}

func Search(b *Bag) bool {
	fmt.Println(b)
	fmt.Println("----")

	if b.Children == nil {
		return false
	}

	for _, child := range b.Children {
		Search(child)
	}
	return true
}
