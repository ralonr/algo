package main

func main() {
	pyramid2('😊', 7)
}

func pyramid(r rune, n int) {
	c := string(r)
	column := (n + n) - 1 // total columns formula
	for i := 1; i < n+1; i++ {
		// total columns = n +n -1
		//total spaces in a row = columns - (row position + row position - 1)
		level := i + i - 1
		spaces := column - level
		sideSpaces := spaces / 2
		for s1 := 0; s1 < sideSpaces; s1++ {
			print("  ")
		}
		for x := 0; x < column-spaces; x++ {
			print(c)
		}
		for s1 := 0; s1 < sideSpaces; s1++ {
			print("  ")
		}
		print("\n")
	}
}

func pyramid2(r rune, n int) {
	c := string(r)
	columns := 2*n - 1
	midpoint := columns / 2
	for row := 0; row < n; row++ {
		for column := 0; column < columns; column++ {
			if midpoint-row <= column && midpoint+row >= column {
				print(c)
			} else {
				print("  ")
			}
		}
		print("\n")
	}
}
