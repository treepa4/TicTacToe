package greeting

import "fmt"

func Greeting() {
	fmt.Println("----------------")
	defer fmt.Println("----------------")
	fmt.Println("Привет! Добро пожаловать в игру \"Крестики - Нолики\". \nПравила ты и так уже знаешь. Размер поля 3x3\nУдачи! ")

}
