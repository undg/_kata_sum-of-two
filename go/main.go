package main

import (
	"errors"
)

var ErrNotFound = errors.New("not found")

// O(n²)
func twoSumSkip(nums []int, target int) ([]int, error) {
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

// O(n)
func twoSum(nums []int, target int) ([]int, error) {
	hashTable := make(map[int]int)
	for i, n := range nums {
		ii, ok := hashTable[target-n]
		if ok {
			return []int{ii, i}, nil
		}

		hashTable[n] = i
	}
	return nil, ErrNotFound
}
