package main

import "fmt"

func main() {
	//test()
	//test2()
	//test3()
	//test4()
	test5()
}

/* Объявление структуры и сразу единичная запись в неё */
func test() {
	employee := struct {
		name   string
		sex    string // пол
		age    int
		salary int // зарплата
	}{
		name:   "Вася",
		sex:    "М",
		age:    25,
		salary: 1500,
	}

	fmt.Printf("%+v\n", employee)
}

type employee struct {
	name   string
	sex    string // пол
	age    int
	salary int // зарплата
}

/* Обычная функция которая возвращает структуру - employee */
func newEmployee(name, sex string, age, salary int) employee {
	return employee{
		name:   name,
		sex:    sex,
		age:    age,
		salary: salary,
	}
}

/*
	Объявление методов структур имеет немного различный синтаксис с функциями, и выглядит он следующим образом:
	func (s structName) methodName(arg1 type, arg2 type ...) returnType {}
	где:
	- structName — имя структуры для которой объявляется метод.
	- s — так называемый ресивер (или получатель, но я больше люблю использовать английскую терминологию).
	Хорошей практикой будет использовать 1–2 символа, которые являются первыми символами названия структуры.
*/
func (e employee) getInfo() string {
	return fmt.Sprintf("Сотрудник: %s\nПол: %s\nВозраст: %d\nЗарплата: %d", e.name, e.sex, e.age, e.salary)
}

/*
	Указатель на структуру.

	Объекты структур также могут быть указателями. В таком случае, передавая наш объект в другие функции по ссылке,
	мы сможем менять значения внутри самого объекта, по аналогии со срезами и мапами.
	Передав в setName() объект структуры по ссылке, мы получили возможность изменять значение полей самого объекта внутри функции.
*/
func setName(e *employee, name string) {
	e.name = name
}

/*
	Указатель в роли ресивера.

	Давайте рассмотрим отдельный тип методов, который в качестве ресивера принимает указатель,
	что позволяет изменять сами поля объекта внутри метода.

	Теперь setName() является методом структуры employee, а в качестве ресивера принимает указатель (e *employee).

	Методы, которые принимают в качестве ресивера значение (без звездочки), работают внутри тела метода с копией объекта.
	Поэтому любые изменения этой копии не скажутся на объекте.
*/
func (e *employee) setName(name string) {
	e.name = name
}

func test2() {
	e := newEmployee("Вася", "М", 11, 1500)
	fmt.Printf("%+v\n", e)
}
func test3() {
	firstEmployee := newEmployee("Вася", "М", 22, 1500)
	fmt.Println(firstEmployee.getInfo())
}
func test4() {
	firstEmployee := newEmployee("Вася", "М", 25, 1500)
	setName(&firstEmployee, "Гена")
	fmt.Println(firstEmployee.getInfo())
}
func test5()  {
	firstEmployee := newEmployee("Вася", "М",25, 15010)
	firstEmployee.setName("Гена2")
	fmt.Println(firstEmployee.getInfo())
}
