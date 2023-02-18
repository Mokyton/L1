package main

import (
	"go.uber.org/goleak"
	"testing"
)

func TestRun(t *testing.T) {
	defer goleak.VerifyNone(t)
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test_1",
			args: args{data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
		},
		{
			name: "Test_1",
			args: args{data: []int{2, 4, 5, 7, 58, 56, 12, 12, 1, 1, 21, 2, 12, 1, 2, 1, 21, 2, 1, 2, 1, 21, 2, 1, 4, 3, 43213, 42, 34, 23, 4, 23, 4, 23, 42, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Run(tt.args.data)
		})
	}
}
