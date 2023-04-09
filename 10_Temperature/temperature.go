package main

func TemperatureGroups(temperatures []float64) map[int][]float64 {
	tempMap := make(map[int][]float64)
	for _, temp := range temperatures {
		step := int(temp) - (int(temp) % 10) // шаг в 10, отнимаем разряд единиц
		tempMap[step] = append(tempMap[step], temp)
	}
	return tempMap
}
