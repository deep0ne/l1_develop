// 10. Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

package main

func TemperatureGroups(temperatures []float64) map[int][]float64 {
	tempMap := make(map[int][]float64) // создаём мапу с ключем в виде шага и значением в виде слайса температур
	for _, temp := range temperatures {
		step := int(temp) - (int(temp) % 10) // шаг в 10, отнимаем разряд единиц
		tempMap[step] = append(tempMap[step], temp)
	}
	return tempMap
}
