package day08

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//
// Create array of arrays
// Get len of first sub item,  plus 2
// get len of sub slice - 2
// loop over iner items
// check current value unitl edge,  on top, bottom , rigth, left
// if one of the sides is visible ,  marks as count

func First() {

	// items := [][]int{
	// 	{3, 0, 3, 7, 3},
	// 	{2, 5, 5, 1, 2},
	// 	{6, 5, 3, 3, 2},
	// 	{3, 3, 5, 4, 9},
	// 	{3, 5, 3, 9, 0},
	// }

	file, err := os.Open("day08/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	items := [][]int{}

	for sc.Scan() {
		currText := sc.Text()
		height := strings.Split(currText, "")

		item := []int{}
		for _, h := range height {
			number, err := strconv.Atoi(h)
			if err != nil {
				panic(err)
			}
			item = append(item, number)
		}
		items = append(items, item)
	}

	visible := 0
	for y := 1; y < len(items)-1; y++ {
		// fmt.Println(items[y])
		for x := 1; x < len(items[y])-1; x++ {
			// fmt.Println(items[y][x])
			current := items[y][x]

			// check toop
			isVisibleTop := true
			top := y - 1
			for top >= 0 {
				if items[top][x] >= current {
					isVisibleTop = false
					break
				}
				top--
			}

			// rigth
			isVisibleRigth := true
			rigth := x + 1
			for rigth <= len(items[y])-1 {
				if items[y][rigth] >= current {
					isVisibleRigth = false
					break
				}
				rigth++
			}

			// bottom
			isVisibleBottom := true
			bottom := y + 1
			for bottom <= len(items[y])-1 {
				if items[bottom][x] >= current {
					isVisibleBottom = false
					break
				}
				bottom++
			}

			// left
			isVisibleLeft := true
			left := x - 1
			for left >= 0 {
				if items[y][left] >= current {
					isVisibleLeft = false
					break
				}
				left--
			}

			if isVisibleTop || isVisibleRigth || isVisibleBottom || isVisibleLeft {
				visible++
			}
		}
	}

	fmt.Println(len(items[0])*2 + (len(items)-2)*2 + visible)
}

func Two() {
	file, err := os.Open("day08/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	items := [][]int{}

	for sc.Scan() {
		currText := sc.Text()
		height := strings.Split(currText, "")

		item := []int{}
		for _, h := range height {
			number, err := strconv.Atoi(h)
			if err != nil {
				panic(err)
			}
			item = append(item, number)
		}
		items = append(items, item)
	}

	scenicScore := []int{}
	for y := 1; y < len(items)-1; y++ {
		for x := 1; x < len(items[y])-1; x++ {
			current := items[y][x]

			// check toop
			treeTop := 0
			top := y - 1
			for top >= 0 {
				treeTop++
				if items[top][x] >= current {
					break
				}
				top--
			}

			// rigth
			treeRigth := 0
			rigth := x + 1
			for rigth <= len(items[y])-1 {
				treeRigth++
				if items[y][rigth] >= current {
					break
				}
				rigth++
			}

			// bottom
			treeBottom := 0
			bottom := y + 1
			for bottom <= len(items[y])-1 {
				treeBottom++
				if items[bottom][x] >= current {
					break
				}
				bottom++
			}

			// left
			treeLeft := 0
			left := x - 1
			for left >= 0 {
				treeLeft++
				if items[y][left] >= current {
					break
				}
				left--
			}

			result := treeTop * treeLeft * treeRigth * treeBottom
			scenicScore = append(scenicScore, result)
		}
	}

	fmt.Println(scenicScore)

	sort.Ints(scenicScore)
	fmt.Println(scenicScore[len(scenicScore)-1])

}
