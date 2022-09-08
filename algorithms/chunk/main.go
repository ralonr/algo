package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	x := chunkSlice(a, 2)
	fmt.Println(x)
}

// chunkSlice re-slice to slice of slice with given size
// if 0 is as size, returning slice will contain s as one element
func chunkSlice(s []int, size int) [][]int {
	chunk := make([][]int, 0)
	if size < 1 {
		chunk = append(chunk, s)
		return chunk
	}
	for i := 0; i < len(s); i += size {
		end := i + size
		if end > len(s) {
			end = len(s)
		}
		chunk = append(chunk, s[i:end])
	}
	return chunk
}
