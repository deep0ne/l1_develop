// 23. Удалить i-ый элемент из слайса.

package main

// Если нам важен порядок (но будет медленнее)
func DeleteFromSlice(nums []int, index int) []int {
	if index < 0 || index > len(nums)-1 {
		return nil
	}
	return append(nums[:index], nums[index+1:]...)
}

// Если порядок не важен
func DeleteFromSliceFaster(nums []int, index int) []int {
	if index < 0 || index > len(nums)-1 {
		return nil
	}
	nums[index] = nums[len(nums)-1]
	return nums[:len(nums)-1]
}
