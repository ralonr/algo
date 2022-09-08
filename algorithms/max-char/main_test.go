package main

import "testing"

//Benchmark_maxChar-4   	  900985	      1188 ns/op
func Benchmark_maxChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = maxChar("xxxxXXXXXXYYYYvvvv")
	}
}

//Benchmark_maxChar2-4   	  910224	      1151 ns/op
func Benchmark_maxChar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = maxChar("xxxxXXXXXXYYYYvvvv")
	}
}
