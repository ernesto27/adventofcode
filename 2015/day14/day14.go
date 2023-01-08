package day14

import (
	"fmt"
	"sort"
)

type Reno struct {
	name               string
	kmBySecond         int
	secondsToRest      int
	secondsToRestCount int
	flyCount           int
	flySeconds         int
	flySecondsCount    int
	points             int
}

var ren1 = &Reno{
	name:               "Rudolph",
	kmBySecond:         22,
	secondsToRest:      165,
	secondsToRestCount: 165,
	flySeconds:         8,
	flySecondsCount:    8,
}
var ren2 = &Reno{
	name:               "Cupid",
	kmBySecond:         8,
	secondsToRest:      114,
	secondsToRestCount: 114,
	flySeconds:         17,
	flySecondsCount:    17,
}

var ren3 = &Reno{
	name:               "Prancer",
	kmBySecond:         18,
	secondsToRest:      103,
	secondsToRestCount: 103,
	flySeconds:         6,
	flySecondsCount:    6,
}

var ren4 = &Reno{
	name:               "Donner",
	kmBySecond:         25,
	secondsToRest:      145,
	secondsToRestCount: 145,
	flySeconds:         6,
	flySecondsCount:    6,
}

var ren5 = &Reno{
	name:               "Dasher",
	kmBySecond:         11,
	secondsToRest:      125,
	secondsToRestCount: 125,
	flySeconds:         12,
	flySecondsCount:    12,
}

var ren6 = &Reno{
	name:               "Comet",
	kmBySecond:         21,
	secondsToRest:      121,
	secondsToRestCount: 121,
	flySeconds:         6,
	flySecondsCount:    6,
}

var ren7 = &Reno{
	name:               "Blitzen",
	kmBySecond:         18,
	secondsToRest:      50,
	secondsToRestCount: 50,
	flySeconds:         3,
	flySecondsCount:    3,
}

var ren8 = &Reno{
	name:               "Vixen",
	kmBySecond:         20,
	secondsToRest:      75,
	secondsToRestCount: 75,
	flySeconds:         4,
	flySecondsCount:    4,
}

var ren9 = &Reno{
	name:               "Dancer",
	kmBySecond:         7,
	secondsToRest:      119,
	secondsToRestCount: 119,
	flySeconds:         20,
	flySecondsCount:    20,
}

func First() {
	renos := []*Reno{ren1, ren2, ren3, ren4, ren5, ren6, ren7, ren8, ren9}

	for x := 1; x <= 2503; x++ {
		for _, ren := range renos {
			if ren.flySecondsCount != 0 {
				ren.flyCount += ren.kmBySecond
				ren.flySecondsCount--
			} else {
				ren.secondsToRestCount--
				if ren.secondsToRestCount == 0 {
					ren.flySecondsCount = ren.flySeconds
					ren.secondsToRestCount = ren.secondsToRest
				}
			}
		}
	}

	greatest := 0

	for _, r := range renos {
		if r.flyCount > greatest {
			greatest = r.flyCount
		}
	}

	fmt.Println(greatest)
}

func Second() {
	renos := []*Reno{ren1, ren2, ren3, ren4, ren5, ren6, ren7, ren8, ren9}

	for x := 1; x <= 2503; x++ {
		for _, ren := range renos {
			if ren.flySecondsCount != 0 {
				ren.flyCount += ren.kmBySecond
				ren.flySecondsCount--
			} else {
				ren.secondsToRestCount--
				if ren.secondsToRestCount == 0 {
					ren.flySecondsCount = ren.flySeconds
					ren.secondsToRestCount = ren.secondsToRest
				}
			}
		}

		// Get renos with the highest points
		sort.Slice(renos, func(i, j int) bool {
			return renos[i].flyCount > renos[j].flyCount
		})

		// Update renos points
		maxValue := 0
		for index, reno := range renos {
			if index == 0 {
				maxValue = reno.flyCount
				reno.points++
				continue
			}
			if reno.flyCount >= maxValue {
				reno.points++
			}
		}
	}

	greatest := 0
	for _, r := range renos {
		if r.points > greatest {
			greatest = r.points
		}
	}
	fmt.Println(greatest)
}
