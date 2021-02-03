package main

import (
	"testing"
)

func TestAddUpper(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"first", args{n: 10}, 10,
		},
		{
			"second", args{n: 20}, 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddUpper(tt.args.n); got != tt.want {
				t.Errorf("AddUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}
