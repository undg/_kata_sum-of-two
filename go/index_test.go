package main

import (
	"reflect"
	"testing"
)

// Import the implementation here.
// Since I don't see a go file yet, I'll assume it will be in the same package or imported.
// For now, I'll define a placeholder for the purpose of creating the test structure.
func twoSum(nums []int, target int) []int {
	// Placeholder implementation - this will be replaced by the actual logic
	return nil
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{1, 2, 3}, 3, []int{0, 1}},
		{[]int{1, 3, 2}, 3, []int{0, 2}},
		{[]int{30, 15, 10, 15, 30}, 30, []int{1, 3}},
	}

	for _, tt := range tests {
		got := twoSum(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("twoSum(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}
