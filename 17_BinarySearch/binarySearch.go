// 17. Реализовать бинарный поиск встроенными методами языка.

package main

// бинарный поиск, который ищет индекс, по которому находится искомый элемент
func binarySearch(numbers []int, num int) int {
	// создаём два указателя в самом начале и в самом конце
	l, r := 0, len(numbers)-1

	// сдвигаем эти два указателя друг к другу в зависимости от условия
	for l <= r {
		// выбираем средний индекс слайса
		mid := l + (r-l)/2

		if numbers[mid] == num { // если нашли, возвращаем
			return mid
		} else if numbers[mid] > num { // если число по индексу больше искомого, значит оно находится в левой половине
			r = mid - 1 // сдвигаем правый указатель, чтобы осталась только левая половина
		} else {
			l = mid + 1 // или же сдвигаем левый, чтобы осталась только правая половина
		}
	}

	return -1 // если число не нашлось, возвращаем -1
}
