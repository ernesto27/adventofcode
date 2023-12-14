package day07

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	// Using iota to create an enumerated sequence
	HighCard = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var cardLabels = make(map[string]int)

type Hand struct {
	cards     string
	bid       int
	typeValue int
	rank      int
}

func (h *Hand) isFiveOfAKind() bool {
	firstChar := h.cards[0]
	for _, c := range h.cards {
		if c != rune(firstChar) {
			return false
		}
	}

	return true
}

func (h *Hand) isFourOfAKind() bool {
	m := make(map[rune]int)

	for _, c := range h.cards {
		m[c]++
	}

	if len(m) != 2 {
		return false
	}

	for _, v := range m {
		if v == 4 {
			return true
		}
	}

	return false
}

func (h *Hand) isFullHouse() bool {
	m := make(map[rune]int)

	for _, c := range h.cards {
		m[c]++
	}

	if len(m) != 2 {
		return false
	}

	for _, v := range m {
		if v == 3 {
			return true
		}
	}

	return false
}

func (h *Hand) isThreeOfAKind() bool {
	m := make(map[rune]int)

	for _, c := range h.cards {
		m[c]++
	}

	if len(m) != 3 {
		return false
	}

	for _, v := range m {
		if v == 3 {
			return true
		}
	}

	return false
}

func (h *Hand) isTwoPair() bool {
	m := make(map[rune]int)

	for _, c := range h.cards {
		m[c]++
	}

	if len(m) != 3 {
		return false
	}

	for _, v := range m {
		if v > 2 {
			return false
		}
	}

	return true
}

func (h *Hand) isOnePair() bool {
	m := make(map[rune]int)

	for _, c := range h.cards {
		m[c]++
	}

	if len(m) != 4 {
		return false
	}

	for _, v := range m {
		if v > 2 {
			return false
		}
	}

	return true
}

func (h *Hand) isGreatherThan(other Hand) bool {
	for i := 0; i < len(h.cards); i++ {
		if cardLabels[string(h.cards[i])] > cardLabels[string(other.cards[i])] {
			return true
		} else if cardLabels[string(h.cards[i])] < cardLabels[string(other.cards[i])] {
			return false
		}
	}
	return false
}

type Hands []Hand

func (p Hands) Len() int           { return len(p) }
func (p Hands) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Hands) Less(i, j int) bool { return p[i].typeValue < p[j].typeValue }

func First() {

	cardLabels["2"] = 1
	cardLabels["3"] = 2
	cardLabels["4"] = 3
	cardLabels["5"] = 4
	cardLabels["6"] = 5
	cardLabels["7"] = 6
	cardLabels["8"] = 7
	cardLabels["9"] = 8
	cardLabels["T"] = 9
	cardLabels["J"] = 10
	cardLabels["Q"] = 11
	cardLabels["K"] = 12
	cardLabels["A"] = 13

	hands := Hands{}

	data, err := os.ReadFile("day07/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	for _, l := range lines {
		a := strings.Split(l, " ")
		//fmt.Println(a)
		n, err := strconv.Atoi(a[1])
		if err != nil {
			panic(err)
		}

		hands = append(hands, Hand{
			cards: a[0],
			bid:   n,
		})

	}

	// hands = append(hands, Hand{
	// 	cards: "32T3K",
	// 	bid:   765,
	// })

	// hands = append(hands, Hand{
	// 	cards: "T55J5",
	// 	bid:   684,
	// })

	// hands = append(hands, Hand{
	// 	cards: "KK677",
	// 	bid:   28,
	// })

	// hands = append(hands, Hand{
	// 	cards: "KTJJT",
	// 	bid:   220,
	// })

	// hands = append(hands, Hand{
	// 	cards: "QQQJA",
	// 	bid:   483,
	// })

	for i, _ := range hands {

		if hands[i].isFiveOfAKind() {
			hands[i].typeValue = FiveOfAKind
			continue
		}

		if hands[i].isFourOfAKind() {
			hands[i].typeValue = FourOfAKind
			continue
		}

		if hands[i].isFullHouse() {
			hands[i].typeValue = FullHouse
			continue
		}

		if hands[i].isThreeOfAKind() {
			hands[i].typeValue = ThreeOfAKind
			continue
		}

		if hands[i].isTwoPair() {
			hands[i].typeValue = TwoPair
			continue
		}

		if hands[i].isOnePair() {
			hands[i].typeValue = OnePair
			continue
		}

		hands[i].typeValue = HighCard
	}

	sort.Sort(hands)
	//fmt.Println(hands)

	for i := 0; i < len(hands); i++ {
		fmt.Println(hands[i].cards)

		// check next repeat
		if i < len(hands)-1 {
			for hands[i+1].typeValue == hands[i].typeValue {
				if hands[i].isGreatherThan(hands[i+1]) {
					hands[i+1].rank = i + 1
					hands[i].rank = i + 2

					temp := hands[i]
					hands[i] = hands[i+1]
					hands[i+1] = temp

				} else {
					hands[i].rank = i + 1
					hands[i+1].rank = i + 2
				}
				i += 1

				if i > len(hands)-1 {
					break
				}
			}

			if i < len(hands)-1 && hands[i].rank == 0 {
				hands[i].rank = i + 1
			}

		} else {
			hands[i].rank = i + 1
		}
	}

	fmt.Println(hands)

	res := 0

	for _, h := range hands {
		res += h.bid * h.rank
	}

	fmt.Println(res)

	// fmt.Println(hands[0].isFiveOfAKind())
	// fmt.Println(hands[4].isFourOfAKind())
	// fmt.Println(hands[4].isFullHouse())
	// fmt.Println(hands[4].isThreeOfAKind())
	// fmt.Println(hands[4].isTwoPair())
	// fmt.Println(hands[4].isOnePair())

}
