package main

import "fmt"

func main() {
	rightSteps('*', 10)
	leftStep('*', 10)
}

func rightSteps(r rune, n int) {
	c := string(r)
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(c)
		}
		fmt.Print("\n")
	}
}

func leftStep(r rune, n int) {
	s := string(r)
	for i := 0; i < n; i++ {
		for j := i; j < n-1; j++ {
			print(" ")
		}
		for x := n - i; x < n+1; x++ {
			print(s)
		}
		print("\n")
	}
}
