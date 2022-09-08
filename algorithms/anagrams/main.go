package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
)

func main() {
	fmt.Println(isAnagram("a b c123 d XYS A B D ! X @", "a b c123 d XYS A XB D ! @"))
	fmt.Println(isAnagram2("a b c123 bd XYS A B D ! X @", "a b c123 d XYS A XB D ! @"))
}

func isAnagram(a, b string) bool {
	ac, err := cleanString(a)
	if err != nil {
		log.Fatal(err)
	}
	bc, err := cleanString(b)
	if err != nil {
		log.Fatal(err)
	}

	if len(ac) != len(bc) {
		return false
	}

	return compareCharMaps(charMap(strings.ToLower(ac)), charMap(strings.ToLower(bc)))
}

func cleanString(s string) (string, error) {
	re, err := regexp.Compile(`[\W]`)
	if err != nil {
		log.Fatal(err)
	}
	s = string(re.ReplaceAll([]byte(s), []byte("")))
	return s, nil
}

func charMap(s string) map[rune]int {
	m := make(map[rune]int, 0)
	for _, v := range []rune(s) {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	return m
}

func compareCharMaps(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if r, ok := b[k]; ok {
			if v != r {
				return false
			}
		}
	}
	return true
}

// method 2

func isAnagram2(a, b string) bool {
	a, _ = cleanString(a)
	b, _ = cleanString(b)

	if len(a) != len(b) {
		return false
	}
	ac := stringToCharSlice(a)
	bc := stringToCharSlice(b)
	sort.Strings(ac)
	sort.Strings(bc)
	for i, av := range ac {
		return av == bc[i]
	}
	return true
}

func stringToCharSlice(s string) []string {
	r := []rune(s)
	sl := make([]string, 0)
	for _, v := range r {
		sl = append(sl, string(v))
	}
	return sl
}
