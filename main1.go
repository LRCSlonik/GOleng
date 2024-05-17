package main

import (
	"fmt"
	"math"
)

func main() {
	//ax^2 + bx + c = 0
	var a float64
	var b float64
	var c float64
	fmt.Println("Реши квадратное уравнение")

	fmt.Println("Введи а:")
	fmt.Scan(&a)
	fmt.Println("Введи b:")
	fmt.Scan(&b)
	fmt.Println("Введи c:")
	fmt.Scan(&c)

	D := (b * b) - 4*(a*c)

	if D > 0 {
		var x1 float64
		var x2 float64

		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)

		fmt.Println("Ваше уровнение имеет 2 корня\n D = " + fmt.Sprint(D))
		fmt.Println("X1: " + fmt.Sprint(x1) + "\nX2: " + fmt.Sprint(x2))
	} else if D == 0 {
		var x float64

		x = (-b) / (2 * a)
		fmt.Println("Ваше уравнение имеет 1 корень\nD = 0")
		fmt.Println("X: " + fmt.Sprint(x))
	} else if D < 0 {
		fmt.Println("Ваше уровнение не имеет корней\nD < 0\nD = " + fmt.Sprint(D))
	}

	/*
		var number float64 = 275.256
		fmt.Println(number)
		age := 16
		fmt.Println(age)
		name := "denis"
		fmt.Println(name)
		fmt.Println(len(name))
	*/

	/*var name string
	var age uint8
	fmt.Println("Как тебя зову?")
	fmt.Scan(&name)
	fmt.Println("Привет " + name + "!!!")
	fmt.Scan(&age)
	fmt.Println("Тебе " + fmt.Sprint(age) + "лет!")


	num := -4
	if num > 0 {
		fmt.Println("Значение больше нуля")
	} else if num < 0 {
		fmt.Println("Значение меньше нуля")
	} else if num == 3 {
		fmt.Println("Значение равно 3 !!!")
	}
	*/
}
