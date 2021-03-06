package main

import (
	"errors"
	"fmt"
)

/*
Пишем собственное хранилище
Допустим, мы разрабатываем очень простую систему хранения данных для учета сотрудников.
Нам нужно уметь сохранять, получать и удалять информацию из этого хранилища по нашим, уже знакомым с предыдущего урока, сотрудникам.ерфейс. Тип соответствует (удовлетворяет) интерфейсу,
если он обладает всеми методами, которые требует интерфейс.
*/

/*
Мы с вами создали переменную типа storage. Поскольку интерфейсы — это такие же кастомные типы, как и структуры,
мы можем создавать переменные данного типа.

После создания пустой переменной типа нашего интерфейса, ее значение равно nil и ее тип также равен nil.
Сначала это может показаться вам немного сложным для понимания.

Концептуально значение интерфейсного типа, или просто значение интерфейса, имеет два компонента — конкретный тип и значение этого типа.
Они называются динамическим типом и динамическим значением интерфейса.

Почему динамическим?
В Go мы можем присваивать в переменную интерфейса другие типы, которые соответствуют данному интерфейсу.
Как только значение интерфейса становится не nil, его динамический тип становится типом нового значения.

Именно поэтому мы без проблемы присвоили в переменную s типа storage объект типа *memoryStorage — этот тип соответствует интерфейсу,
потому что обладает всеми необходимыми методами.

И поэтому после инициализации интерфейса его значение и тип равны nil, а после присвоения нашей структуры, значение не равно nil,
а тип равен *memoryStorage.

Если вы переименуете метод delete() на remove(), то компилятор начнет ругаться, а редактор подчеркнет ошибку.

Как только мы изменили имя метода, наша структура memoryStorage уже не реализует интерфейс storage.
По этому присвоение данного объекта в переменную типа интерфейса стало невозможным.
*/
func main() {
	var s storage
	fmt.Println("s == nil", s == nil)
	fmt.Printf("type of s: %T\n", s)

	d := newMemoryStorage()
	fmt.Println("d == nil", d == nil)
	fmt.Printf("type of d: %T\n", d)
}

type employee struct {
	id int
	name string
	age int
	salary int
}

type storage interface {
	insert(e employee) error
	get(id int) (employee, error)
	delete(id int) error
}

type memoryStorage struct {
	data map[int]employee
}

/*
	Описшем конструктор, который возвращает указатель на структуру и инициализирует мапу при создании объекта.
	Если этого не сделать, мы получим ошибку при попытке записи в неинициализированную мапу.
*/
func newMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]employee),
	}
}

/*
	Реализация методов интерфейса
*/

//Указатель на структуру memoryStorage который реализует функцию insert, которая является методом структуры employee
func (s *memoryStorage) insert(e employee) error {
	s.data[e.id] = e
	return nil
}

func (s *memoryStorage) get(id int) (employee, error) {
	e, exists := s.data[id]
	if !exists {
		return employee{}, errors.New("employee with such id doesn't exist")
	}

	return e, nil
}

func (s *memoryStorage) delete(id int) error {
	delete(s.data, id)
	return nil
}