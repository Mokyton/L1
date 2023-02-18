package main

import (
	"go.uber.org/goleak"
	"testing"
)

func TestRun1(t *testing.T) {
	defer goleak.VerifyNone(t)
	tests := []struct {
		name string
	}{
		{""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Run1()
		})
	}
}

func TestRun2(t *testing.T) {
	defer goleak.VerifyNone(t)
	tests := []struct {
		name string
	}{
		{""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Run2()
		})
	}
}

func TestRun3(t *testing.T) {
	defer goleak.VerifyNone(t)
	tests := []struct {
		name string
	}{
		{""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Run3()
		})
	}
}

func TestRun4(t *testing.T) {
	defer goleak.VerifyNone(t)
	tests := []struct {
		name string
	}{
		{""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Run4()
		})
	}
}

func TestRun5(t *testing.T) {
	defer goleak.VerifyNone(t)
	tests := []struct {
		name string
	}{
		{""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Run5()
		})
	}
}
