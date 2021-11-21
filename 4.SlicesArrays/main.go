package main

import "fmt"

func main() {
	sliceLinkToArray()
}

func sliceCapacity() {
	var todoList = []string{
		"купить хлеб",
		"купить молоко",
		"купить пиво",
		"сделать уборку",
	}

	newTodoList := append(todoList, "купить девушке цветы")

	fmt.Printf("Длина=%d и емкость=%d у старого среза\n", len(todoList), cap(todoList))
	fmt.Printf("Длина=%d и емкость=%d у нового среза\n", len(newTodoList), cap(newTodoList))
}

func sliceLinkToArray() {
	todoList := [4]string{
		"купить хлеб",
		"купить молоко",
		"купить пиво",
		"дописать третий модуль",
	}

	var tasks []string
	fmt.Println(tasks == nil)

	tasks = todoList[1:4]
	fmt.Println(tasks == nil)

	for i := range tasks {
		fmt.Println(tasks[i])
	}
}