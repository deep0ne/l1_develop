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
	xDistance := math.Pow(float64(p2.x-p1.x), 2)
	yDistance := math.Pow(float64(p2.y-p1.y), 2)
	return math.Sqrt(xDistance + yDistance)
}
