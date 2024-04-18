package main

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple test #1",
			args: args{n: 3, k: 2},
			want: 3,
		},
		{
			name: "simple test #2",
			args: args{n: 5, k: 2},
			want: 8,
		},
		{
			name: "simple test #3",
			args: args{n: 3, k: 3},
			want: 4,
		},
		{
			name: "simple test #4",
			args: args{n: 5, k: 3},
			want: 13,
		},
		{
			name: "when n is a big number",
			args: args{n: 1000000, k: 3},
			want: 5926420018876522305,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_climbStairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		climbStairs(10000000, 10)
	}
}

// Benchmark_climbStairs-12             5   209340612 ns/op  80003110 B/op        1 allocs/op
