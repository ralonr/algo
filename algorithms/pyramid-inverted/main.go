package main

import (
	"fmt"
)

func main() {
	invertedPyramid('😊', 5)
}

func invertedPyramid(r rune, n int) {
	// total columns = n +n -1
	// total spaces in a row = columns - (row position + row position - 1)
	cols := n + n - 1
	for i := n; i >= 0; i-- {
		c := string(r)
		sps := cols - (i + i - 1)
		for sp := 0; sp < (sps / 2); sp++ {
			fmt.Print("  ")
		}

		for x := 0; x < cols-sps; x++ {
			print(c)
		}

		for sp := 0; sp < (sps / 2); sp++ {
			fmt.Print("  ")
		}
		fmt.Print("\n")
	}
}
