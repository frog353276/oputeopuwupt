package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
		{10, -5, 5},
	}

	for _, test := range tests {
		result := Add(test.a, test.b)
		if result != test.expected {
			t.Errorf("Add(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{5, 3, 2},
		{0, 0, 0},
		{1, -1, 2},
		{10, 15, -5},
	}

	for _, test := range tests {
		result := Subtract(test.a, test.b)
		if result != test.expected {
			t.Errorf("Subtract(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 6},
		{0, 5, 0},
		{-2, 3, -6},
		{-2, -3, 6},
	}

	for _, test := range tests {
		result := Multiply(test.a, test.b)
		if result != test.expected {
			t.Errorf("Multiply(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b, expected int
		shouldError    bool
	}{
		{6, 2, 3, false},
		{10, 5, 2, false},
		{7, 3, 2, false}, // integer division
		{5, 0, 0, true},  // division by zero
	}

	for _, test := range tests {
		result, err := Divide(test.a, test.b)

		if test.shouldError {
			if err == nil {
				t.Errorf("Divide(%d, %d) should have returned an error", test.a, test.b)
			}
		} else {
			if err != nil {
				t.Errorf("Divide(%d, %d) returned unexpected error: %v", test.a, test.b, err)
			}
			if result != test.expected {
				t.Errorf("Divide(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
			}
		}
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		{2, true},
		{3, false},
		{0, true},
		{-2, true},
		{-3, false},
	}

	for _, test := range tests {
		result := IsEven(test.n)
		if result != test.expected {
			t.Errorf("IsEven(%d) = %t; expected %t", test.n, result, test.expected)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{3, 5, 5},
		{5, 3, 5},
		{-1, -5, -1},
		{0, 0, 0},
	}

	for _, test := range tests {
		result := Max(test.a, test.b)
		if result != test.expected {
			t.Errorf("Max(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		n, expected int
	}{
		{5, 5},
		{-5, 5},
		{0, 0},
		{-1, 1},
	}

	for _, test := range tests {
		result := Abs(test.n)
		if result != test.expected {
			t.Errorf("Abs(%d) = %d; expected %d", test.n, result, test.expected)
		}
	}
}
