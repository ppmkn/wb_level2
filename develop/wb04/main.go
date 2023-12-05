package main

import (
	"sort"
	"strings"
)

// FindAnagrams находит все множества анаграмм в словаре
func FindAnagrams(words []string) map[string][]string {
	// Инициализируем карту для хранения множеств анаграмм
	anagramSets := make(map[string][]string)

	// Проходим по каждому слову в словаре
	for _, word := range words {
		// Приводим слово к нижнему регистру
		word = strings.ToLower(word)

		// Сортируем руны слова для создания уникального ключа
		key := sortRunes(word)

		// Если ключ уже существует в карте, добавляем текущее слово в соответствующее множество
		anagramSets[key] = append(anagramSets[key], word)
	}

	// Удаляем множества из одного элемента
	for key, set := range anagramSets {
		if len(set) < 2 {
			delete(anagramSets, key)
		} else {
			// Сортируем множество по возрастанию
			sort.Strings(set)
		}
	}

	return anagramSets
}

// sortRunes сортирует руны строки и возвращает результат как строку
func sortRunes(s string) string {
	// Преобразуем строку в слайс рун
	runes := []rune(s)
	// Сортируем руны
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	// Преобразуем отсортированный слайс рун обратно в строку
	return string(runes)
}

func main() {
	// Пример использования
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagramSets := FindAnagrams(words)

	// Вывод результатов
	for key, set := range anagramSets {
		println("Множество анаграмм для ключа", key, ":", strings.Join(set, ", "))
	}
}
