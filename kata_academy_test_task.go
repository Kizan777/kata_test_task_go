package main

import (
	"bufio"
	"fmt"
	"os"
	sl "slices"
	"strconv"
	st "strings"
)

func check_input(arr, arabic_digits, roman_digits, operations []string) string {
	if len(arr) != 3 {
		panic("Неверное количество аргументов")
	}
	is_first_arabic := sl.Contains(arabic_digits, arr[0])
	is_first_roman := sl.Contains(roman_digits, arr[0])
	is_second_arabic := sl.Contains(arabic_digits, arr[2])
	is_second_roman := sl.Contains(roman_digits, arr[2])
	is_operation := sl.Contains(operations, arr[1])
	if is_first_arabic && is_second_arabic && is_operation {
		return "арабские"
	} else if is_first_roman && is_second_roman && is_operation {
		return "римские"
	} else if is_first_arabic && is_second_roman || is_first_roman && is_second_arabic {
		panic("Используются одновременные разные системы счисления")
	} else {
		panic("Введены числа в некорректных диапазонах или дробные числа")
	}
}

func roman_to_arabic(roman_num string, arabic_digits, roman_digits []string) string {
	return arabic_digits[sl.Index(roman_digits, roman_num)]
}

func arabic_to_roman(arabic_num int, arabic_digits, roman_digits, roman_tens []string) string {
	if arabic_num < 10 {
		return roman_digits[sl.Index(arabic_digits, strconv.Itoa(arabic_num))]
	} else {
		if arabic_num%10 == 0 {
			return roman_tens[sl.Index(arabic_digits, strconv.Itoa(arabic_num/10))]
		} else {
			return roman_tens[sl.Index(arabic_digits, strconv.Itoa(arabic_num/10))] +
				roman_digits[sl.Index(arabic_digits, strconv.Itoa(arabic_num%10))]
		}
	}
}

func arabic_calculations(arr []string) int {
	num1, _ := strconv.Atoi(arr[0])
	num2, _ := strconv.Atoi(arr[2])
	switch arr[1] {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		panic("Некорректная операция")
	}
}

func roman_calculations(arr, arabic_digits, roman_digits, roman_tens []string) string {
	num1 := roman_to_arabic(arr[0], arabic_digits, roman_digits)
	num2 := roman_to_arabic(arr[2], arabic_digits, roman_digits)
	answer := arabic_calculations([]string{num1, arr[1], num2})
	if answer < 1 {
		panic("Отрицательный или нулевой ответ для римских чисел")
	} else {
		return arabic_to_roman(answer, arabic_digits, roman_digits, roman_tens)
	}
}

func main() {
	arabic_digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	roman_digits := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	roman_tens := []string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC", "C"}
	operations := []string{"+", "-", "*", "/"}
	reader := bufio.NewReader(os.Stdin)
	for {
		msg, _ := reader.ReadString('\n')
		msg = st.TrimSpace(msg)
		arr := st.Split(msg, " ")
		//Проверяем корректность ввода
		if check_input(arr, arabic_digits, roman_digits, operations) == "арабские" {
			//работа с арабскими
			fmt.Println(arabic_calculations(arr))
		} else {
			//работа с римскими
			fmt.Println(roman_calculations(arr, arabic_digits, roman_digits, roman_tens))
		}
		break
	}

}
