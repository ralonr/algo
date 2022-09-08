package main

import "fmt"

func main() {
	fmt.Println(fib(8))
	fmt.Println(fib(10))
	fmt.Println(fib(26))
}

func fib2(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func fib(n int) int {
	a := []int{0, 1}

	for i := 2; i <= n; i++ {
		a = append(a, a[i-1]+a[i-2])
	}
	return a[len(a)-1]
}
