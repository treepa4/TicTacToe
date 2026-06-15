package greeting

import "fmt"

func Greeting() {
	fmt.Println("----------------")
	defer fmt.Println("----------------")
	fmt.Println("Привет! Добро пожаловать в игру \"Крестики - Нолики\". \nПравила ты и так уже знаешь. Размер поля 3x3\nУдачи! ")
	fmt.Println("Игровое поле:")
	for i := 0; i < 3; i++ {
		fmt.Printf("|_|_|_|\n")
	}
}
