package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CutOptions хранит параметры для утилиты cut
type CutOptions struct {
	fields    string
	delimiter string
	separated bool
}

// cut обрабатывает входные данные в соответствии с параметрами cutOptions
func cut(input string, cutOptions CutOptions) string {
	var result strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()

		// Проверяем, содержит ли строка разделитель
		if strings.Contains(line, cutOptions.delimiter) || !cutOptions.separated {
			// Разбиваем строку на поля
			fields := strings.Split(line, cutOptions.delimiter)

			// Выбираем запрошенные поля
			for _, fieldNumStr := range strings.Split(cutOptions.fields, ",") {
				fieldNum := fieldNumStrAsInt(fieldNumStr)
				if fieldNum > 0 && fieldNum <= len(fields) {
					result.WriteString(fields[fieldNum-1])
					result.WriteString("\t")
				}
			}

			// Создаем новый strings.Builder для каждой строки
			lineResult := result.String()
			result.Reset()

			// Убираем лишний таб в конце строки
			if len(lineResult) > 0 {
				lineResult = lineResult[:len(lineResult)-1]
			}

			result.WriteString(lineResult)
			result.WriteString("\n")
		}
	}

	return result.String()
}

// fieldNumStrAsInt конвертирует строку в число с учетом обработки ошибок
func fieldNumStrAsInt(fieldNumStr string) int {
	fieldNum, err := strconv.Atoi(fieldNumStr)
	if err != nil {
		return 0
	}
	return fieldNum
}

func main() {
	// Парсим флаги командной строки
	fields := flag.String("f", "", "выбрать поля (колонки)")
	delimiter := flag.String("d", "\t", "использовать другой разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	// Инициализируем CutOptions
	cutOptions := CutOptions{
		fields:    *fields,
		delimiter: *delimiter,
		separated: *separated,
	}

	// Читаем данные из STDIN
	scanner := bufio.NewScanner(os.Stdin)
	var input strings.Builder
	for scanner.Scan() {
		input.WriteString(scanner.Text())
		input.WriteString("\n")
	}

	// Вызываем функцию cut и выводим результат
	output := cut(input.String(), cutOptions)
	fmt.Print(output)
}
