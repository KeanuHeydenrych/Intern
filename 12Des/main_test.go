package main

import (
	"testing"
)

func TestMySum(t *testing.T) {
	x := mySum(2, 3)
	if mySum(2, 3) != 5 {
		t.Error("Error, the sum is WRONG. Expected:", 5, "Got:", x)
	}

}

func TestMySum2(t *testing.T) {
	x := mySum2(2, 3)
	if mySum2(2, 3) != 5 {
		t.Error("Error, the sum is WRONG. Expected:", 5, "Got:", x)
	}

}

// Table tests
func TestMySum3(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{68, 1}, 69},
		test{[]int{67, 2}, 69},
		test{[]int{52, 17}, 69},
		test{[]int{1, 2, 3, 63}, 69},
	}

	for _, v := range tests {
		x := mySum3(v.data...)
		if x != v.answer {
			t.Error("Expected:", v.answer, "Got:", x)
		}

	}

}
