package main

import (
	"errors"
	"strings"
	"unicode"
)

// Unpack выполняет примитивную распаковку строки
func Unpack(s string) (string, error) {
	var result strings.Builder // Создаем переменную result типа strings.Builder для построения строки
	var repeatCount int        // Переменная для хранения количества повторений символа
	var escaped bool           // Флаг, указывающий, что предыдущий символ был '\'

	for _, char := range s {
		if escaped {
			if unicode.IsDigit(char) {
				repeatCount = repeatCount*10 + int(char-'0') // Если символ после '\' - цифра, добавляем к числу повторений
			} else {
				// Обработка escape-последовательности
				result.WriteString(strings.Repeat(string(char), repeatCount)) // Повторяем символ нужное количество раз и добавляем в результат
				repeatCount = 0                                               // Сбрасываем счетчик повторений
				escaped = false                                               // Сбрасываем флаг escape
			}
		} else {
			if unicode.IsDigit(char) {
				return "", errors.New("некорректная строка") // Если цифра в обычной части строки, считаем строку некорректной
			} else if char == '\\' {
				escaped = true // Если символ '\', устанавливаем флаг escape
			} else {
				// Обычный символ
				if repeatCount == 0 {
					repeatCount = 1 // Если не указано количество повторений, считаем, что символ повторяется один раз
				}
				result.WriteString(strings.Repeat(string(char), repeatCount)) // Повторяем символ нужное количество раз и добавляем в результат
				repeatCount = 0                                               // Сбрасываем счетчик повторений
			}
		}
	}

	if escaped {
		return "", errors.New("некорректная строка") // Если строка заканчивается на '\', считаем строку некорректной
	}

	return result.String(), nil // Возвращаем строку, полученную после распаковки, и отсутствие ошибки
}

func main() {
	// Пример использования
	result, err := Unpack("a4bc2d5e")
	if err != nil {
		println("Ошибка:", err.Error())
	} else {
		println(result) // Ожидаемый вывод: "aaaabccddddde"
	}
}