package main

import "fmt"

func main() {
	var chunks []int

	chunkLength := 3
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range iter(nums) {
		if (i+1)%chunkLength == 0 {
			chunks = nums[i+1-chunkLength : i+1]
			fmt.Printf("chunk: %v\n", chunks)
			continue
		}

		if (i + 1) == len(nums) {
			n := len(nums) / chunkLength
			chunks = nums[n*chunkLength:]
			fmt.Printf("chunk: %v\n", chunks)
			break
		}
	}
}

func iter(nums []int) func(yield func(i int) bool) bool {
	return func(yield func(i int) bool) bool {
		for i := range nums {
			if !yield(i) {
				return false
			}
		}
		return true
	}
}
