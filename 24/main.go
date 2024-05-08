// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
package main

import (
	"fmt"
	"math"
)

// структура точки с координатами
type point struct {
	x, y float64
}

// Конструктор
func newPoint(x, y float64) *point {
	return &point{
		x: x,
		y: y,
	}
}

// findDistance находит расстояние между двумя точками
func (p *point) findDistance(p2 *point) float64 {
	var distance float64
	// формула нахождения d = √((хА – хВ)^2 + (уА – уВ)^2)
	ox, oy := p.x-p2.x, p.y-p2.y
	distance = math.Sqrt(math.Pow(ox, 2) + math.Pow(oy, 2))
	return distance
}

func main() {
	p1 := newPoint(10, 20)
	p2 := newPoint(30, 60)

	distance := p1.findDistance(p2)
	fmt.Println(distance)
}
