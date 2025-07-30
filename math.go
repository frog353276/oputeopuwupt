package main

import "errors"

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func IsEven(n int) bool {
	return n%2 == 0
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
}
