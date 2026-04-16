package main

import (
	"errors"
)

var ErrNotFound = errors.New("not found")

// O(n*n)
func twoSum(nums []int, target int) ([]int, error) {
	for i, n := range nums {
		tailIndexStart := i + 1
		tail := nums[tailIndexStart:]
		for ii, nn := range tail {
			if target == n+nn {
				return []int{i, tailIndexStart + ii}, nil
			}
		}
	}
	return nil, ErrNotFound
}
