package main

func DeleteFromSlice(nums []int, index int) []int {
	if index < 0 || index > len(nums)-1 {
		return nil
	}
	return append(nums[:index], nums[index+1:]...)
}
