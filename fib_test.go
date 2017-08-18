package main

import (
	"testing"
)

func TestFib(t *testing.T) {
	results := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765}
	for i, n := range results {
		if f := Fib(i); n != f {
			t.Error("Expected", n, "Produced", f)
		}
	}
}
