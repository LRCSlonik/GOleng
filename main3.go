package main

import "fmt"

func main() {
	fmt.Println("Строковый калькулятор")
	fmt.Println("Введите выражение")

	var a string
	var b string
	var c string
	fmt.Scan(&a, &b, &c)

	//if b != +
	//fmt.Println("Ошибка ввода используйте (+, -, *, /)")

	switch {
	case b == "+":
		fmt.Println(a + c)
	case b == "-":
		fmt.Println(a - c)

	}
}
