package main

import "fmt"

func maxSubArray(nums []int) int {

	sum := 0
	maxSum := nums[0]

	for i := 0; i < len(nums); i++ {

		sum = sum + nums[i]

		if sum > maxSum {
			maxSum = sum
		}

		if sum < 0 {
			sum = 0
		}

	}

	return maxSum

}

func main() {

	var arr = []int{-2}

	fmt.Println(maxSubArray(arr))

}
