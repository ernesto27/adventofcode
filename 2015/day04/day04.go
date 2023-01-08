package day04

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func First() {
	key := "ckczppom"

	i := 0
	for {
		s := strconv.Itoa(i)
		hasher := md5.New()
		hasher.Write([]byte(key + s))
		h := hex.EncodeToString(hasher.Sum(nil))

		isValid := false
		for x := 0; x < 6; x++ {
			if string(h[x]) == "0" {
				isValid = true
			} else {
				isValid = false
				break
			}
		}

		if isValid {
			fmt.Println(s)
			break
		}

		i++
	}
}
