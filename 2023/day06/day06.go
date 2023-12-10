package day06

import "fmt"

type Race struct {
	time     int
	distance int
}

func First() {
	races := []Race{}
	// races = append(races, Race{
	// 	time:     7,
	// 	distance: 9,
	// })
	// races = append(races, Race{
	// 	time:     15,
	// 	distance: 40,
	// })
	// races = append(races, Race{
	// 	time:     30,
	// 	distance: 200,
	// })

	races = append(races, Race{
		time:     44,
		distance: 208,
	})
	races = append(races, Race{
		time:     80,
		distance: 1581,
	})
	races = append(races, Race{
		time:     65,
		distance: 1050,
	})
	races = append(races, Race{
		time:     72,
		distance: 1102,
	})

	res := 0

	for _, r := range races {
		count := 0

		for i := 1; i <= r.time; i++ {
			m := i * (r.time - i)
			if m > r.distance {
				count++
			}
		}

		if count > 0 {
			if res == 0 {
				res = count
			} else {
				res *= count
			}
		}

	}

	fmt.Println(res)
}

func Second() {

	time := 44806572
	distance := 208158110501102
	count := 0
	res := 0

	for i := 1; i <= time; i++ {
		m := i * (time - i)
		if m > distance {
			count++
		}
	}

	if count > 0 {
		if res == 0 {
			res = count
		} else {
			res *= count
		}
	}

	fmt.Println(res)
}
