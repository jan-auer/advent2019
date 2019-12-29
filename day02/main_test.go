package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestReset(t *testing.T) {
	input := Memory{1, 0, 0, 0, 99}
	program := NewProgram(input)

	program.Run()
	program.Reset()
	actual := program.mem

	if !reflect.DeepEqual(input, actual) {
		t.Errorf("Reset(%v) => Got %v", input, actual)
	}
}

func TestRerun(t *testing.T) {
	input := Memory{1, 0, 0, 0, 99}
	program := NewProgram(input)

	program.Run()
	run1 := program.mem

	program.Reset()
	program.Run()
	run2 := program.mem

	if !reflect.DeepEqual(run1, run2) {
		t.Errorf("Run(%v) => Got %v, expected %v", input, run2, run1)
	}
}

func TestRun(t *testing.T) {
	var tests = []struct {
		input, expected Memory
	}{
		{
			Memory{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			Memory{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},

		{
			Memory{1, 0, 0, 0, 99},
			Memory{2, 0, 0, 0, 99},
		},

		{
			Memory{2, 3, 0, 3, 99},
			Memory{2, 3, 0, 6, 99},
		},

		{
			Memory{2, 4, 4, 5, 99, 0},
			Memory{2, 4, 4, 5, 99, 9801},
		},

		{
			Memory{1, 1, 1, 4, 99, 5, 6, 0, 99},
			Memory{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	for _, tt := range tests {
		program := NewProgram(tt.input)
		program.Run()
		actual := program.mem
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("Run(%v) => Got %v, want %v", tt.input, actual, tt.expected)
		}
	}
}

func TestRead(t *testing.T) {
	var tests = []struct {
		input    string
		expected Memory
	}{
		{
			"1,0,0,0,99",
			Memory{1, 0, 0, 0, 99},
		},

		{
			"99",
			Memory{99},
		},
	}

	for _, tt := range tests {
		actual, err := ReadMemory(strings.NewReader(tt.input))
		if err != nil {
			t.Errorf("Read(%v) => Error reading: %v", tt.input, err)
		}

		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("Read(%q) => Got %v, want %v", tt.input, actual, tt.expected)
		}
	}
}
