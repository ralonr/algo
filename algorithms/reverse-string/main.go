package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse("abcdef"))
	fmt.Println(reverse("xyz"))
	fmt.Println(reverse("123456789"))
	fmt.Println(reverse("123 456 789"))
}

func reverse(s string) string {
	ns := []rune(s)
	for i, j := 0, len(ns)-1; i < j; i, j = i+1, j-1 {
		ns[i], ns[j] = ns[j], ns[i]
	}
	return string(ns)
}
