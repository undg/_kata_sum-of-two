package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		target  int
		want    []int
		wantErr error
	}{
		{"simple", []int{1, 2, 3}, 3, []int{0, 1}, nil},
		{"simple reorder", []int{1, 3, 2}, 3, []int{0, 2}, nil},
		{"bigger numbers", []int{30, 15, 10, 15, 30}, 30, []int{1, 3}, nil},
		{"not found", []int{30, 14, 10, 15, 30}, 30, nil, ErrNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) || !errors.Is(err, tt.wantErr) {
				t.Errorf("twoSum(%v, %d) = %v; want %v, %v", tt.nums, tt.target, got, tt.want, tt.wantErr)
			}
		})
	}
}
