package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// grepOptions содержит опции командной строки для утилиты grep
type grepOptions struct {
	afterContext  int      // -A: количество строк после совпадения
	beforeContext int      // -B: количество строк перед совпадением
	context       int      // -C: общее количество строк вокруг совпадения (A + B)
	count         bool     // -c: выводить только количество строк
	ignoreCase    bool     // -i: игнорировать регистр
	invert        bool     // -v: исключать строки совпадения
	fixedString   string   // -F: точное совпадение строки, не паттерн
	lineNumber    bool     // -n: выводить номера строк
	pattern       string   // Паттерн поиска
	files         []string // Файлы для поиска
}

// grepResult содержит результат поиска для одного файла
type grepResult struct {
	fileName string   // Имя файла
	lineNum  int      // Номер строки
	line     string   // Строка совпадения
	after    []string // Строки после совпадения
	before   []string // Строки перед совпадением
}

func main() {
	// Обработка командной строки и получение опций grep
	options := parseCommandLine()

	// Выполнение поиска в каждом файле и вывод результатов
	for _, file := range options.files {
		results := grepFile(file, options)
		printResults(results, options)
	}
}

// parseCommandLine разбирает опции командной строки и возвращает структуру grepOptions
func parseCommandLine() grepOptions {
	// Инициализация структуры опций
	var options grepOptions

	// Задание значений по умолчанию
	flag.IntVar(&options.afterContext, "A", 0, "печатать N строк после совпадения")
	flag.IntVar(&options.beforeContext, "B", 0, "печатать N строк перед совпадением")
	flag.IntVar(&options.context, "C", 0, "печатать ±N строк вокруг совпадения (A + B)")
	flag.BoolVar(&options.count, "c", false, "вывести только количество строк")
	flag.BoolVar(&options.ignoreCase, "i", false, "игнорировать регистр")
	flag.BoolVar(&options.invert, "v", false, "исключать строки совпадения")
	flag.StringVar(&options.fixedString, "F", "", "точное совпадение строки, не паттерн")
	flag.BoolVar(&options.lineNumber, "n", false, "выводить номера строк")
	flag.Parse()

	// Обработка паттерна поиска (позиционный аргумент)
	options.pattern = flag.Arg(0)

	// Обработка файлов для поиска (позиционные аргументы)
	options.files = flag.Args()[1:]

	return options
}

// grepFile выполняет поиск в одном файле согласно опциям grep
func grepFile(fileName string, options grepOptions) []grepResult {
	// Открытие файла
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка открытия файла:", err)
		return nil
	}
	defer file.Close()

	// Инициализация результата поиска
	var results []grepResult

	// Сканер для чтения файла по строкам
	scanner := bufio.NewScanner(file)

	// Регулярное выражение для поиска
	var regex *regexp.Regexp
	if options.fixedString != "" {
		// Если используется точное совпадение строки, создаем регулярное выражение
		regex = regexp.MustCompile(regexp.QuoteMeta(options.fixedString))
	} else {
		// Иначе создаем регулярное выражение на основе паттерна
		regex = regexp.MustCompile(options.pattern)
	}

	// Флаги для контроля состояния поиска
	found := false
	matchCount := 0
	afterCount := options.afterContext
	beforeCount := options.beforeContext
	afterLines := make([]string, 0, options.afterContext)
	beforeLines := make([]string, 0, options.beforeContext)

	// Обработка каждой строки файла
	for lineNum := 1; scanner.Scan(); lineNum++ {
		line := scanner.Text()

		// Проверка на соответствие паттерну
		isMatch := regex.MatchString(line)

		// Учет игнорирования регистра
		if options.ignoreCase {
			isMatch = regex.MatchString(strings.ToLower(line))
		}

		// Учет инвертированного поиска
		if options.invert {
			isMatch = !isMatch
		}

		// Обработка найденного совпадения
		if isMatch {
			found = true
			matchCount++

			// Сохранение строк после совпадения
			afterLines = append(afterLines, line)
			afterCount = options.afterContext

			// Учет строк перед совпадением
			if options.beforeContext > 0 {
				beforeLines = append(beforeLines, line)
				// Обрезаем слайс, чтобы поддерживать только beforeContext строк перед совпадением
				if len(beforeLines) > options.beforeContext {
					beforeLines = beforeLines[1:]
				}
			}
		} else {
			// Обработка строк после совпадения
			if found && afterCount > 0 {
				afterLines = append(afterLines, line)
				afterCount--
			}

			// Обработка строк перед совпадением
			if beforeCount > 0 {
				beforeLines = append(beforeLines, line)
				beforeCount--
			}
		}

		// Сброс флага found, если достигнуто количество строк после совпадения
		if found && afterCount == 0 {
			found = false
		}

		// Добавление результата поиска в список, если найдено совпадение
		if found && options.count {
			continue
		}

		if found && !options.count {
			result := grepResult{
				fileName: fileName,
				lineNum:  lineNum,
				line:     line,
				after:    append([]string{}, afterLines...), // Копирование для изоляции состояния
				before:   append([]string{}, beforeLines...),
			}
			results = append(results, result)
		}
	}

	return results
}

// printResults выводит результаты поиска в соответствии с опциями
func printResults(results []grepResult, options grepOptions) {
	// Обработка результата поиска для каждого файла
	for _, result := range results {
		// Вывод номера файла, если ищем в нескольких файлах
		if len(options.files) > 1 {
			fmt.Printf("%s:", result.fileName)
		}

		// Вывод номера строки, если включена опция -n
		if options.lineNumber {
			fmt.Printf("%d:", result.lineNum)
		}

		// Вывод строки совпадения
		fmt.Println(result.line)

		// Вывод строк перед совпадением
		for _, beforeLine := range result.before {
			fmt.Println(beforeLine)
		}

		// Вывод строк после совпадения
		for _, afterLine := range result.after {
			fmt.Println(afterLine)
		}
	}
}
