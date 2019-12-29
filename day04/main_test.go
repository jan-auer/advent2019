package main

import (
	"testing"
)

func TestCross(t *testing.T) {
	var tests = []struct {
		check func(password) bool
		value int
		ok    bool
	}{
		{checkPart1, 111111, true},
		{checkPart1, 223450, false},
		{checkPart1, 123789, false},
		{checkPart2, 112233, true},
		{checkPart2, 123444, false},
		{checkPart2, 111122, true},
	}

	for _, tt := range tests {
		ok := tt.check(NewPassword(tt.value))
		if tt.ok != ok {
			t.Errorf("check(%v) => Got %v, want %v", tt.value, ok, tt.ok)
		}
	}
}
