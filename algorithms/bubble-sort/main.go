package main

import "fmt"

func main() {
	fmt.Println(bubbleSort([]int{5, 2, 3, 4, 1}))
}

func bubbleSort(slice []int) ([]int, int) {
	count := 0
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			if slice[i] == slice[j] {
				// mean same element
				break
			}

			if slice[i] < slice[j] {
				slice[j], slice[i] = slice[i], slice[j]
				count++
			}
		}
	}
	return slice, count
}
