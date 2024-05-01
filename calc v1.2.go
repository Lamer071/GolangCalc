package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RomanToArabic(roman string) int {
	romanMap := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}

	return romanMap[roman]
}

func Calculate(operation string) int {
	tokens := strings.Split(operation, "+")
	if len(tokens) != 2 {
		panic("Неверный формат математической операции")
	}

	a, b := tokens[0], tokens[1]

	var result int

	if a >= "1" && a <= "10" && b >= "1" && b <= "10" {
		num1, _ := strconv.Atoi(a)
		num2, _ := strconv.Atoi(b)
		result = num1 + num2
	} else if a >= "I" && a <= "X" && b >= "I" && b <= "X" {
		num1 := RomanToArabic(a)
		num2 := RomanToArabic(b)
		result = num1 + num2
	} else {
		panic("Использованы разные системы счисления")
	}

	return result
}

func main() {
	fmt.Println("Введите математическую операцию в формате '1+1' или 'V+V':")
	var operation string
	_, err := fmt.Scanln(&operation)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result := Calculate(operation)
	fmt.Println("Результат:", result)
}
