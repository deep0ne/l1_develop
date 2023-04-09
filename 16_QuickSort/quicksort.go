// 16. Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

// функция для перестановки элементов с учётом опроного
func partition(nums []int, left, right int) int {
	pivot := nums[right] // выбираем опорный элемент
	index := left
	for left <= right-1 { // переставляем таким образом, чтобы все элементы меньше опорного оказались слева, а больше - справа
		if nums[left] < pivot {
			nums[left], nums[index] = nums[index], nums[left]
			index++
		}
		left++
	}
	nums[index], nums[right] = nums[right], nums[index]
	return index
}

// рекурсивно разбиваем и сортируем слайс
func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	part := partition(nums, left, right)
	// применяем к двум подмассивам - левому и правому
	quickSort(nums, left, part-1)
	quickSort(nums, part+1, right)
}

// функция для инициализации с двумя указателями
func startQuickSort(nums []int) []int {
	l, r := 0, len(nums)-1
	quickSort(nums, l, r)
	return nums
}
