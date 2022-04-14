package main

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{1, 3}, 3},
		{"test2", args{5, 3}, 5},
		{"test3", args{3, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
func BenchmarkMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		max := Max(7, 3)
		b.Log(max)
	}
}
func ExampleMax() {
	fmt.Println(Max(1, 3))
	// Output:
	// 3
}
