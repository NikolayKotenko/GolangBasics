package main

import "fmt"

func main() {
	mapa4()
}
func mapa1() {
	/* Инициализируем мапу методом - make */
	ages := make(map[string]int)
	ages["Максим"] = 20
	ages["Олег"] = 25
	ages["Саня"] = 28

	for key, value := range ages {
		fmt.Printf("%s - %d\n", key, value)
	}
}

func mapa2() {
	/* Можно инициализировать мапу так */
	ages := map[string]int{
		"Максим": 20,
		"Олег": 25,
		"Саня": 28,
	}

	fmt.Printf("Максиму %d лет\n", ages["Максим"])
}


/**
Чтобы проверить, содержится ли данный ключ в нашей мапе, Go может возвращать 2 значения
по обращению к элементу мапы по ключу: само значение и bool, который true, если такой элемент есть и наоборот.
*/
func mapa3() {
	ages := map[string]int{
		"Максим": 20,
		"Олег": 25,
		"Саня": 28,
	}

	age, exists := ages["Антон"]
	if !exists {
		fmt.Println("Антона нет в списке")
		return
	}
	if age != 0 {
		fmt.Println("age != 0")
		return
	}

	fmt.Printf("Антону %d лет\n", ages["Антон"])
}

/*
Как мы уже говорили, мапы, как и срезы, являются указателями на область в памяти.
Поэтому если скопировать мапу в новую переменную и удалить из нее элементы, это отобразится также и на новой мапе.

Из этого так же выплывает и то, что мапы, как и срезы, при передаче в качестве аргумента функции, передаются по ссылке.
Это значит, что любые изменения внутри тела функции также изменят саму мапу.
*/
func mapa4() {
	ages := map[string]int{
		"Максим": 20,
		"Олег":   25,
		"Саня":   28,
		"Антон":  35,
	}

	var agesNew map[string]int

	fmt.Println(ages)
	fmt.Println(agesNew)

	agesNew = ages

	delete(ages, "Антон")

	fmt.Println(ages)
	fmt.Println(agesNew)
}