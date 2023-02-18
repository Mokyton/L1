package main

import (
	"go.uber.org/goleak"
	"testing"
)

func TestRun(t *testing.T) {
	defer goleak.VerifyNone(t)
	tests := []struct {
		name string
	}{
		{name: "leak test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Run()
		})
	}
}
