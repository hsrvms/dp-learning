package main

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "edge case #1",
			args: args{n: 0},
			want: 0,
		},
		{
			name: "edge case #2",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "edge case #3",
			args: args{n: 2},
			want: 2,
		},
		{
			name: "simple case #1",
			args: args{n: 3},
			want: 3,
		},
		{
			name: "simple case #2",
			args: args{n: 5},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_climbStairs10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		climbStairs(10)
	}
}
