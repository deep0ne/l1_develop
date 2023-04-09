// 24. Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

package main

import "math"

type Point struct {
	x, y float32
}

func NewPoint(x, y float32) Point {
	return Point{
		x: x,
		y: y,
	}
}

func PointDistance(p1, p2 Point) float64 {
	dx := math.Pow(float64(p2.x-p1.x), 2)
	dy := math.Pow(float64(p2.y-p1.y), 2)
	return math.Sqrt(dx + dy)
}
