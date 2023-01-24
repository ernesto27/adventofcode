package day05

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func A() {

	doorID := "ugkcyxxp"
	count := 0
	sum := 0
	var password [8]string

	for count < 8 {
		hasher := md5.New()
		sumStr := strconv.Itoa(sum)
		hasher.Write([]byte(doorID + sumStr))
		h := hex.EncodeToString(hasher.Sum(nil))
		sixDigits := h[:5]

		sum++

		if sixDigits == "00000" {
			// check if 6 char is number
			n, err := strconv.Atoi(string(h[5]))
			if err != nil {
				continue
			}
			if n > 7 {
				continue
			}

			if password[n] == "" {
				password[n] = string(h[6])
				count++

			}

		}
	}

	fmt.Println(password)
}
