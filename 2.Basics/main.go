package main

import "fmt"

func main() {
	wat()
}

func wat() {
	var i, j, k int                 // int, int, int
	var b, f, s = true, 2.3, "four" // bool, float64, string
	fmt.Println(i, j, k)
	fmt.Println(b, f, s)
}

func dataTypes() {
	var number int = -1
	var number2 uint = 2
	var number3 float32 = 3.14
	var isGoCool bool = true
	var name string = "Maksim"

	fmt.Println(number)
	fmt.Println(number2)
	fmt.Println(number3)
	fmt.Println(isGoCool)
	fmt.Println(name)
}

const pi = 3.1415

func calcCircle() {
	circleRadius := 2 //радиус круга в см
	// площадь круга
	circleArea := float32(circleRadius) * float32(circleRadius) * pi

	fmt.Printf("Радиус круга: %d см\n", circleRadius)
	fmt.Println("Формула для расчета площади круга: S=πr2")

	fmt.Printf("Площадь круга: %.2f см. кв\n", circleArea)
}

/*
	Дело в том, что поскольку Go является строго типизированным ЯП, всевозможные математические операции возможны только над переменными одного и того же типа.
	При инициализации circleRadius , посмотрев на значение 2, компилятор определил тип этой переменной как int. Но константа pi является числом с плавающей точкой, поэтому ее компилятор определил как float32.
	Поскольку у этих двух переменных разные типы, мы не можем производить над ними базовые операции. В таких случаях используется приведение типов.
	Мы приводим int к float32 с помощью данного синтаксиса. Также можно делать противоположное приведение (float32 к int). Тогда значение округляется в меньшую сторону.
*/

// Animal has a Name and an Age to represent an animal.
type Animal struct {
	Name string
	Age  uint
}

// String makes Animal satisfy the Stringer interface.
func (a Animal) String() string {
	return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}

func callAnimal() {
	a := Animal{
		Name: "Gopher",
		Age:  2,
	}
	fmt.Println(a)
}
