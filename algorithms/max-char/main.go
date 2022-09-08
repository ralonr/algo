package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(maxChar("abbbbaccccccvvvvvvvvvvv"))
	fmt.Println(maxChar("111222    "))
	fmt.Println(maxChar("111222111"))
	fmt.Println(maxChar("DDDDddddd"))
	fmt.Println(maxChar2("DDDDddddd"))
	fmt.Println(maxChar2("DDDDdddddxxx"))
	fmt.Println(maxChar2("DDDDdddddXXXXXX"))
}

func maxChar(s string) string {
	r := []rune(s)
	max := ""
	maxCount := 0
	for _, v := range r {
		count := 0
		for _, n := range r {
			if v == n {
				count++
			}
		}
		if maxCount < count {
			maxCount = count
			max = string(v)
		}
	}
	return max
}

func maxChar2(s string) string {
	max := ""
	count := 0
	var lastChar int32
	for _, v := range s {
		if lastChar == v { // make it efficient to avoid duplicate char checks
			continue
		}
		c := strings.Count(s, string(v))
		if count < c {
			count = c
			max = string(v)
			lastChar = v
		}
	}
	return max
}
