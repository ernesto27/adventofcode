package day04

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type kv struct {
	Key   string
	Value int
}
type room struct {
	letters  []kv
	sectorID int
	Checksum string
}

func createRoom(input string) room {
	r := strings.Split(input, "-")
	letters := r[:len(r)-1]
	valueCheck := r[len(r)-1:]
	valueCheck = strings.Split(valueCheck[0], "[")
	sectorID, _ := strconv.Atoi(valueCheck[0])
	checksum1 := valueCheck[1][:len(valueCheck[1])-1]
	m := make(map[string]int)

	for _, letter := range letters {
		for _, c := range letter {
			m[string(c)]++
		}
	}

	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Value == ss[j].Value {
			return ss[i].Key < ss[j].Key
		}

		return ss[i].Value > ss[j].Value
	})

	return room{
		letters:  ss,
		sectorID: sectorID,
		Checksum: checksum1,
	}

}

func A() {

	file, _ := ioutil.ReadFile("day04/input.txt")

	input := strings.Split(string(file), "\n")

	var rooms []room
	for _, i := range input {
		rooms = append(rooms, createRoom(i))
	}

	sum := 0
	for _, r := range rooms {
		isValid := true
		for i, c := range r.Checksum {
			if r.letters[i].Key != string(c) {
				isValid = false
				break
			}
		}
		if isValid {
			sum += r.sectorID
		}
	}

	fmt.Println(sum)
}

func B() {
	file, _ := ioutil.ReadFile("day04/input.txt")

	input := strings.Split(string(file), "\n")

	for _, i := range input {
		room := createRoom(i)
		r := strings.Split(i, "-")
		letters := r[:len(r)-1]

		for _, l := range letters {
			for _, c := range l {
				caesar := (((int(c) - 'a') + room.sectorID) % 26) + 'a'
				fmt.Print(string(caesar))
			}
		}
		fmt.Println(room.sectorID)
		fmt.Println("-----")

	}
}
