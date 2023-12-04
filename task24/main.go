package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором
*/

// структура точки
type Point struct {
	x, y int
}

// функция расстояния
func (p *Point) Distance(q *Point) float64 {
	return math.Sqrt(math.Pow(float64(q.x-p.x), 2) + math.Pow(float64(q.y-p.y), 2))
}

// конструктор
func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

func main() {
	p := NewPoint(0, 0)
	q := NewPoint(5, 5)
	fmt.Println(p.Distance(q))
}
