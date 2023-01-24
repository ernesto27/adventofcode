package day13

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type bus struct {
	ID        int
	timestamp int
	diff      int
}

func A() {
	earlyTimestap := 1007125
	buses := []int{
		13, 41, 569, 29, 19, 23, 937, 37, 17,
	}

	resp := bus{
		ID:        0,
		timestamp: math.MaxInt,
	}
	for _, b := range buses {
		for x := 0; x <= math.MaxInt; x += b {
			if x >= earlyTimestap {
				fmt.Println(x)
				if x < resp.timestamp {
					resp.ID = b
					resp.timestamp = x
				}
				break
			}
		}
	}

	fmt.Println((resp.timestamp - earlyTimestap) * resp.ID)
}

func B() {

	input, err := os.Open("day13/raw.txt")
	if err != nil {
		fmt.Println(err, "\nYou need a file called input.txt in this directory")
		return
	}
	defer input.Close()

	// --- START ---
	sc := bufio.NewScanner(input)
	var (
		buses     []int
		timestamp int = 100000000000000
		step      int = 1
	)
	sc.Scan()
	sc.Scan()
	for _, bus := range strings.Split(sc.Text(), ",") {
		timestamp, _ := strconv.Atoi(bus)
		buses = append(buses, timestamp)
	}

	for i, bus := range buses {
		if bus == 0 {
			continue
		}
		for (timestamp+i)%bus != 0 {
			timestamp += step
		}
		step *= bus
	}
	fmt.Println(timestamp)
}
