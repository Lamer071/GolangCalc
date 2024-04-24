package main

import (
	. "fmt"
	"strconv"
	"strings"
)

func romanToArabic(roman string) int {
	romanMap := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	return romanMap[roman]
}

func arabicToRoman(number int) string {
	if number <= 0 || number > 10 {
		panic("Римские числа поддерживаются только от I до X")
	}

	romanMap := map[int]string{
		1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
		6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
	}
	return romanMap[number]
}

func calculate(a int, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("Деление на ноль недопустимо")
		}
		return a / b
	default:
		panic("Неподдерживаемая операция")
	}
}

func main() {
	var input string
	Println("Введите выражение (например, 3 + 2):")
	Scanln(&input)

	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Неверный формат ввода")
	}

	aStr, bStr := parts[0], parts[2]
	operator := parts[1]

	var a, b int
	var result int

	if numA, err := strconv.Atoi(aStr); err == nil {
		a = numA
	} else {
		a = romanToArabic(aStr)
	}

	if numB, err := strconv.Atoi(bStr); err == nil {
		b = numB
	} else {
		b = romanToArabic(bStr)
	}

	result = calculate(a, b, operator)

	if aStr == bStr {
		Printf("Результат: %d\n", result)
	} else {
		Printf("Результат: %s\n", arabicToRoman(result))
	}
}
