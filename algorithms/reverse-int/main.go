package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(reverseInt(-1000))
	fmt.Println(reverseInt(-123456))
	fmt.Println(reverseInt(7654321))
	fmt.Println(reverseInt(-7654321))
	fmt.Println(reverseInt(-20))
	fmt.Println(reverseInt(120))
	fmt.Println(reverseInt(0))
	fmt.Println(reverseInt2(12345))
	fmt.Println(reverseInt2(-12345))
	fmt.Println(reverseInt2(0))
	fmt.Println(reverseInt2(-1000))
	fmt.Println(reverseInt2(1200))
}

// reverseInt reverses the int with string based solution
func reverseInt(n int) int {
	var ns []rune
	var str string
	isNeg := n < 0
	if isNeg {
		n = n * -1
		str = strconv.Itoa(n)
		ns = []rune(str)
	} else {
		str = strconv.Itoa(n)
		ns = []rune(str)
	}

	for i, j := 0, len(ns)-1; i < j; i, j = i+1, j-1 {
		ns[i], ns[j] = ns[j], ns[i]
	}
	rv := string(ns) //reversed string
	c, _ := strconv.ParseInt(rv, 10, 32)
	if isNeg {
		c = c * -1
	}
	return int(c)
}

// reverseInt2reverseInt2 reverses the int with a different approach
func reverseInt2(n int) int { //123
	isNeg := n < 0
	if isNeg {
		n = n * -1
	}
	var num int
	for n > 0 { // 123                  // 12				// 1
		t := n % 10      // 3					// 2				// 0
		n = n / 10       // 12.3 ignores 3 so 12 //1.2 =1			// 0.1 so it ignores =0
		num = num*10 + t // 0 + 3				//3 + 3*10 +2= 32   // 32*10 + 1 =321
	}

	if isNeg {
		return num * -1
	}
	return num
}
