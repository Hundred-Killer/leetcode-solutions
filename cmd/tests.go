package main

func buildArray(nums []int) []int {
	newArray := []int{}

	for i := range nums {
		newArray = append(newArray, nums[nums[i]])
	}

	return newArray
}
