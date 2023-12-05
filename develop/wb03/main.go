package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var (
	columnIndex  int
	numericSort  bool
	reverseSort  bool
	uniqueSort   bool
	monthSort    bool
	ignoreSpaces bool
	checkSorted  bool
	humanSort    bool
)

// Функция init() вызывается перед выполнением main() и инициализирует флаги
func init() {
	flag.IntVar(&columnIndex, "k", 0, "Указание колонки для сортировки")
	flag.BoolVar(&numericSort, "n", false, "Сортировать по числовому значению")
	flag.BoolVar(&reverseSort, "r", false, "Сортировать в обратном порядке")
	flag.BoolVar(&uniqueSort, "u", false, "Не выводить повторяющиеся строки")
	flag.BoolVar(&monthSort, "M", false, "Сортировать по названию месяца")
	flag.BoolVar(&ignoreSpaces, "b", false, "Игнорировать хвостовые пробелы")
	flag.BoolVar(&checkSorted, "c", false, "Проверять отсортированы ли данные")
	flag.BoolVar(&humanSort, "h", false, "Сортировать по числовому значению с учетом суффиксов")
}

func main() {
	// Парсим флаги командной строки
	flag.Parse()

	// Проверяем, что передан только один аргумент (имя файла)
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Использование: sortutil <filename>")
		os.Exit(1)
	}

	// Имя файла, содержащего несортированные строки
	fileName := args[0]

	// Читаем строки из файла
	lines, err := readLines(fileName)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		os.Exit(1)
	}

	// Применяем сортировку в соответствии с переданными флагами
	sortLines(lines)

	// Выводим результат в зависимости от наличия флага проверки отсортированности
	if checkSorted {
		if isSorted(lines) {
			fmt.Println("Данные отсортированы.")
		} else {
			fmt.Println("Данные не отсортированы.")
		}
	} else {
		// Выводим отсортированные строки
		for _, line := range lines {
			fmt.Println(line)
		}
	}
}

// readLines читает строки из файла и возвращает их в виде слайса
func readLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// isSorted проверяет, отсортирован ли слайс строк
func isSorted(lines []string) bool {
	for i := 1; i < len(lines); i++ {
		if lines[i-1] > lines[i] {
			return false
		}
	}
	return true
}

// sortLines выполняет сортировку строк в соответствии с переданными флагами
func sortLines(lines []string) {
	sort.SliceStable(lines, func(i, j int) bool {
		if ignoreSpaces {
			lines[i] = strings.TrimSpace(lines[i])
			lines[j] = strings.TrimSpace(lines[j])
		}

		if humanSort {
			return humanLess(lines[i], lines[j])
		}

		if numericSort {
			return compareNumeric(lines[i], lines[j])
		}

		if monthSort {
			return compareMonths(lines[i], lines[j])
		}

		return lines[i] < lines[j]
	})

	if reverseSort {
		reverse(lines)
	}

	if uniqueSort {
		lines = unique(lines)
	}
}

// compareNumeric сравнивает строки как числа
func compareNumeric(a, b string) bool {
	numA, errA := strconv.Atoi(a)
	numB, errB := strconv.Atoi(b)

	if errA == nil && errB == nil {
		return numA < numB
	}

	return a < b
}

// compareMonths сравнивает строки как месяцы
func compareMonths(a, b string) bool {
	timeA, errA := time.Parse("January", a)
	timeB, errB := time.Parse("January", b)

	if errA == nil && errB == nil {
		return timeA.Before(timeB)
	}

	return a < b
}

// humanLess реализует "человеческое" сравнение строк, учитывая числовые суффиксы
func humanLess(a, b string) bool {
	var (
		suffixA string
		suffixB string
	)

	// Разделяем строку на основную часть и суффикс
	if unicode.IsDigit(rune(a[len(a)-1])) {
		// Находим индекс первого нечислового символа с конца строки
		for i := len(a) - 1; i >= 0; i-- {
			if !unicode.IsDigit(rune(a[i])) {
				suffixA = a[i+1:]
				a = a[:i+1]
				break
			}
		}
	}

	if unicode.IsDigit(rune(b[len(b)-1])) {
		// Находим индекс первого нечислового символа с конца строки
		for i := len(b) - 1; i >= 0; i-- {
			if !unicode.IsDigit(rune(b[i])) {
				suffixB = b[i+1:]
				b = b[:i+1]
				break
			}
		}
	}

	// Сравниваем основные части строк
	if a == b {
		// Если основные части равны, сравниваем числовые суффиксы
		numA, errA := strconv.Atoi(suffixA)
		numB, errB := strconv.Atoi(suffixB)

		if errA == nil && errB == nil {
			return numA < numB
		}
	}

	return a < b
}

// reverse инвертирует порядок строк в слайсе
func reverse(lines []string) {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
}

// unique удаляет дубликаты из слайса строк
func unique(lines []string) []string {
	seen := make(map[string]bool)
	uniqueLines := make([]string, 0, len(lines))

	for _, line := range lines {
		if !seen[line] {
			seen[line] = true
			uniqueLines = append(uniqueLines, line)
		}
	}

	return uniqueLines
}
