package main

import (
	"errors"
	"fmt"
)

func main() {
	initStructure()
}
type contact struct {
	firstName string
	lastName string
	phoneNumber string
	email string
	address string
	dateOfBirth string
}
/* Инициализируем структуру */
func initStructure()  {
	c1 := contact{
		firstName: "Ярик же",
		lastName: "Соляной",
		phoneNumber: "+7445456548",
		email: "yarik_huyarit2000@mail.ru",
		address: "Модных пролетариатов 37а",
		dateOfBirth: "08.11.2000",
	}

	c1.printInfo()
}
/* Реализация метода у структуры ! имеет доступ ко всем полям структуры */
func (c contact) printInfo()  {
	fmt.Printf("Имя: %s\nФамилия: %s\nНомер телефона: %s\nE-mail: %s\nАдресс: %s\nДата рождения: %s\n",
		c.firstName, c.lastName, c.phoneNumber, c.email, c.address, c.dateOfBirth)
}

/* ------------------------ */
func mapje()  {
	contactList := map[string]string{
		"Ярик": "получается, хуярит",
		"Леха": "остался демоном",
		"Колян": "ахуевает с Голэнга",
	}

	//delete(contactList, "Колян")
	for name, message := range contactList{
		fmt.Printf("%s: %s\n", name, message)
	}
}
func arraysOrSlices()  {
	contactList := []string{"Ярик", "Леха", "Колян"}

	fmt.Printf("\nКрасивый for\n")
	for index, value := range contactList{
		fmt.Printf("%d. %s\n", index+1, value)
	}

	contactList = append(contactList, "Я ебу как классно", "Жора")

	fmt.Printf("\nДеолтфный for из Си\n")
	for i := 0; i < len(contactList); i++ {
		fmt.Printf("%s\n", contactList[i])
	}
}

/* ------------------------ */
const winePrice int = 100
func marketWine()  {
	result, err := buyWine(17, 120)
	if err != nil {
		fmt.Println("Покупка неуспешна: ", err.Error())
	} else {
		fmt.Printf("Ваша сдача - %d. Хорошего дня!", result)
	}
}
func buyWine(age, moneyAmount int) (int, error)  {
	if age >=18 && moneyAmount >= winePrice {
		return moneyAmount - winePrice, nil
	} else if age < 18 {
		return moneyAmount, errors.New("только вишневый сок, сынок")
	} else {
		return moneyAmount, errors.New("бабок не ма у тебя, сынок")
	}
}
