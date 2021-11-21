package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}
type Square struct {
	sideLength float32
}
type Circle struct {
	radius float32
}

func (c Square) Area() float32  {
	return c.sideLength * c.sideLength
}
func (c Circle) Area() float32 {
	return float32(math.Pi) * (c.radius * c.radius)
}

/*
	Реализует интерфейс Shape который реализует функцию Area()
	Свобода замены одного типа - другим, называется "взаимозаменяемостью"
*/
func printShapeArea(s Shape)  {
	fmt.Printf("Площадь фигуры: %.2f см\n", s.Area())
}
func main() {
	square := Square{10.5}
	circle := Circle{5.7}
	printShapeArea(square)
	printShapeArea(circle)
}


/*
	Пустой интерфейс не имеет методов.
	поэтому любой тип в Go реализовывает пустой интерфейс. Пример тому функция fmt.Printf()
	Пустой интерфейс не описывает поведение.
*/
