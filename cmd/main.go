package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("3 =", Solution(3))
	fmt.Println("1000 =", Solution(1000))
	fmt.Println("1001 =", Solution(1001))
	fmt.Println("1832 =", Solution(1832))
	fmt.Println("3888 =", Solution(3888))
	fmt.Println("1444 =", Solution(1444))
	fmt.Println("1999 =", Solution(1999))
}

func Solution(number int) (res string) {
	var dict = map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}
	k := 1000
	d := 0
	for k > 0 {
		d = number / k % 10
		switch {
		case d%5 == 4:
			res += dict[k] + dict[(d/5+1)*5*k]
		case d >= 5:
			res += dict[5*k]
			if d > 5 {
				res += strings.Repeat(dict[k], d-5)
			}
		case d <= 3:
			res += strings.Repeat(dict[k], d)
		}
		k /= 10
	}
	return
}
