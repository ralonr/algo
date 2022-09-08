package main

import (
	"reflect"
	"testing"
)

func Test_chunkSlice(t *testing.T) {
	type args struct {
		s    []int
		size int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Success 1",
			args: args{s: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, size: 2},
			want: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9}},
		},

		{
			name: "Success 2",
			args: args{s: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, size: 3},
			want: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			name: "Success 3",
			args: args{s: []int{-11, -22, -33, -44, 5, 6, 7, 8, 9}, size: 3},
			want: [][]int{{-11, -22, -33}, {-44, 5, 6}, {7, 8, 9}},
		},
		{
			name: "Success 4 empty slice",
			args: args{s: []int{}, size: 3},
			want: [][]int{},
		},
		{
			name: "Success 5 with zero size",
			args: args{s: []int{-11, -22, -33, -44, 5, 6, 7, 8, 9}, size: 0},
			want: [][]int{{-11, -22, -33, -44, 5, 6, 7, 8, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chunkSlice(tt.args.s, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chunkSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_chunkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = chunkSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)
	}
}
