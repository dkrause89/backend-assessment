package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		A, B []int
		C    []int
		want int
	}{
		{
			name: "example from problem",
			A:    []int{1, 4, 5, 8},
			B:    []int{4, 5, 9, 10},
			C:    []int{4, 6, 7, 10, 2},
			want: 4,
		},
		{
			name: "single plank single nail",
			A:    []int{1},
			B:    []int{5},
			C:    []int{3},
			want: 1,
		},
		{
			name: "single plank nail outside",
			A:    []int{1},
			B:    []int{5},
			C:    []int{10},
			want: -1,
		},
		{
			name: "one nail nails all",
			A:    []int{1, 2, 3},
			B:    []int{10, 10, 10},
			C:    []int{5},
			want: 1,
		},
		{
			name: "impossible",
			A:    []int{1, 5},
			B:    []int{2, 6},
			C:    []int{10, 11},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Solution(tt.A, tt.B, tt.C)
			if got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
