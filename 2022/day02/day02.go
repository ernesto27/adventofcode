package day02

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Rock = iota + 1
	Paper
	Sissors
)

type player struct {
	selected string
}

type Rules struct {
	Options    map[string]int
	PointsWin  int
	PointsDraw int
	PointsLost int
}
type Game struct {
	Player1 player
	Player2 player
	Rules   Rules
}

func (g Game) GetScorePlayer2() int {
	if g.Rules.Options[g.Player1.selected] == Rock && g.Rules.Options[g.Player2.selected] == Paper {
		return Paper + g.Rules.PointsWin
	}

	if g.Rules.Options[g.Player1.selected] == Sissors && g.Rules.Options[g.Player2.selected] == Paper {
		return Paper + g.Rules.PointsLost
	}

	if g.Rules.Options[g.Player1.selected] == Sissors && g.Rules.Options[g.Player2.selected] == Rock {
		return Rock + g.Rules.PointsWin
	}

	if g.Rules.Options[g.Player1.selected] == Paper && g.Rules.Options[g.Player2.selected] == Rock {
		return Rock + g.Rules.PointsLost
	}

	if g.Rules.Options[g.Player1.selected] == Paper && g.Rules.Options[g.Player2.selected] == Sissors {
		return Sissors + g.Rules.PointsWin
	}

	if g.Rules.Options[g.Player1.selected] == Rock && g.Rules.Options[g.Player2.selected] == Sissors {
		return Sissors + g.Rules.PointsLost
	}

	return g.Rules.Options[g.Player2.selected] + g.Rules.PointsDraw
}

func (g Game) GetScorePlayerPartTwo() int {
	if g.Rules.Options[g.Player2.selected] == Paper {
		g.Player2.selected = g.Player1.selected
	} else if g.Rules.Options[g.Player2.selected] == Rock {
		if g.Rules.Options[g.Player1.selected] == Rock {
			g.Player2.selected = "Z"
		} else if g.Rules.Options[g.Player1.selected] == Paper {
			g.Player2.selected = "X"
		} else if g.Rules.Options[g.Player1.selected] == Sissors {
			g.Player2.selected = "Y"
		}

	} else if g.Rules.Options[g.Player2.selected] == Sissors {
		if g.Rules.Options[g.Player1.selected] == Rock {
			g.Player2.selected = "Y"
		} else if g.Rules.Options[g.Player1.selected] == Sissors {
			g.Player2.selected = "X"
		}
	}

	if g.Rules.Options[g.Player1.selected] == Rock && g.Rules.Options[g.Player2.selected] == Paper {
		return Paper + g.Rules.PointsWin
	}

	if g.Rules.Options[g.Player1.selected] == Sissors && g.Rules.Options[g.Player2.selected] == Paper {
		return Paper + g.Rules.PointsLost
	}

	if g.Rules.Options[g.Player1.selected] == Sissors && g.Rules.Options[g.Player2.selected] == Rock {
		return Rock + g.Rules.PointsWin
	}

	if g.Rules.Options[g.Player1.selected] == Paper && g.Rules.Options[g.Player2.selected] == Rock {
		return Rock + g.Rules.PointsLost
	}

	if g.Rules.Options[g.Player1.selected] == Paper && g.Rules.Options[g.Player2.selected] == Sissors {
		return Sissors + g.Rules.PointsWin
	}

	if g.Rules.Options[g.Player1.selected] == Rock && g.Rules.Options[g.Player2.selected] == Sissors {
		return Sissors + g.Rules.PointsLost
	}

	return g.Rules.Options[g.Player2.selected] + g.Rules.PointsDraw
}

func First() {
	rules := Rules{
		Options:    make(map[string]int),
		PointsWin:  6,
		PointsDraw: 3,
		PointsLost: 0,
	}
	rules.Options["A"] = Rock
	rules.Options["B"] = Paper
	rules.Options["C"] = Sissors

	rules.Options["X"] = Rock
	rules.Options["Y"] = Paper
	rules.Options["Z"] = Sissors

	file, _ := os.Open("day02/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	result := 0
	for sc.Scan() {
		currText := sc.Text()
		items := strings.Split(currText, " ")
		game := Game{
			Player1: player{
				selected: items[0],
			},
			Player2: player{
				selected: items[1],
			},
			Rules: rules,
		}

		result += game.GetScorePlayer2()
	}

	fmt.Println(result)
}

func Two() {
	rules := Rules{
		Options:    make(map[string]int),
		PointsWin:  6,
		PointsDraw: 3,
		PointsLost: 0,
	}
	rules.Options["A"] = Rock
	rules.Options["B"] = Paper
	rules.Options["C"] = Sissors

	rules.Options["X"] = Rock
	rules.Options["Y"] = Paper
	rules.Options["Z"] = Sissors

	file, _ := os.Open("day02/input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	result := 0
	for sc.Scan() {
		currText := sc.Text()
		items := strings.Split(currText, " ")
		game := Game{
			Player1: player{
				selected: items[0],
			},
			Player2: player{
				selected: items[1],
			},
			Rules: rules,
		}
		score := game.GetScorePlayerPartTwo()
		fmt.Println(score)
		result += score
	}

	fmt.Println(result)
}
