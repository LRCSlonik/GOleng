package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Добро пожаловать! Введите  выражение:\n")

	intType, first, second, sign, err := readLine()
	if err != nil {
		fmt.Println("Возникла ошибка:\n", err)
		return
	}
	if intType == "arab" {
		firstNum, err1 := strconv.Atoi(first)
		if err1 != nil {
			fmt.Println("Возникла ошибка:\n", err1)
			return
		}
		secondNum, err2 := strconv.Atoi(second)
		if err2 != nil {
			fmt.Println("Возникла ошибка:\n", err2)
			return
		}
		res, err3 := calculator(firstNum, secondNum, sign)
		if err3 != nil {
			fmt.Println("Возникла ошибка:\n", err3)
			return
		} else {
			fmt.Println("Ответ: ", res)
		}
	} else {
		firstNum := fromRomanToInt(first)
		secondNum := fromRomanToInt(second)
		res, err1 := calculator(firstNum, secondNum, sign)
		if err1 != nil {
			fmt.Println("Возникла ошибка при работе калькулятора:\n", err1)
			return
		} else {
			final, err2 := fromIntToRoman(res)
			if err2 != nil {
				fmt.Println("Возникла ошибка при работе калькулятора:\n", err2)
				return
			}
			fmt.Println("Ответ: ", final)
		}
	}
}
