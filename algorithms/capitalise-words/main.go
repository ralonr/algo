package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(capitalise("a text is a \"text\"       always"))
	fmt.Println(capitalise2("a text is a text always"))
}

func capitalise(s string) string {
	c := strings.ToUpper(s[:1])
	for i := 1; i < len(s); i++ {
		if s[i-1:i] == " " {
			c += strings.ToUpper(s[i : i+1])
		} else {
			c += s[i : i+1]
		}
	}
	return c
}

// if more spaces between it will panic
func capitalise2(s string) string {
	a := strings.Split(s, " ")
	for i, v := range a {
		a[i] = strings.ToUpper(v[0:1]) + v[1:]
	}
	return strings.Join(a, " ")
}
