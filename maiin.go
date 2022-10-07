package main

import "fmt"

func main() {

	nums := []int{99}
	target := 2

	fmt.Println(search(nums, target))

}

func search(nums []int, target int) int {
	lower := 0
	upper := len(nums) - 1
	mid := 0
	guess := 0

	for lower <= upper {
		mid = (lower + upper) / 2
		guess = nums[mid]

		if guess == target {
			return mid
		}

		if guess > target {
			upper = mid - 1
		} else {
			lower = mid + 1
		}
	}
	return -1
}
