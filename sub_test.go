package main

import "testing"

func TestSub(t *testing.T) {
	testCases := []struct {
		a, b     int
		expected int
	}{
		{4, 2, 2},
		{-6, 3, -9},
		{0, 5, -5},
		{-2, -2, 0},
		{7, -3, 10},
	}

	for _, tc := range testCases {
		result := Sub(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Sub(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
		}
	}

}
