package main

import (
	"fmt"
	"testing"
)

func TestGuassion(t *testing.T) {
	// [-1, 1]  [0, 1],  [1, 1]
	// [-1, 0]  [0, 0],  [0, 1]
	// [-1, -1] [0, -1], [1, -1]
	arrs := [][2]int{
		{-1, 1}, {0, 1}, {1, 1},
		{-1, 0}, {0, 0}, {0, 1},
		{-1, -1}, {0, -1}, {1, -1},
	}
	for i, arr := range arrs {
		t.Logf("result %d = %v", i, gaussion(arr[0], arr[1], 1.5))
	}
}

func TestGuassionMatrics(t *testing.T) {
	matrics := gaussioniMatrix(1, 1.5)
	for _, arr := range matrics {
		for _, v := range arr {
			fmt.Printf("%v ", v)
		}
		fmt.Printf("\n")
	}
}
