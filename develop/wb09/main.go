package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	// Парсинг флагов
	url := flag.String("url", "", "URL для загрузки")
	output := flag.String("output", "", "Название файла для сохранения")
	flag.Parse()

	// Проверка обязательного флага -url
	if *url == "" {
		fmt.Println("Необходимо указать URL с помощью флага -url")
		return
	}

	// Извлечение имени файла из URL, если -output не указан
	if *output == "" {
		*output = path.Base(*url)
	}

	// Загрузка файла
	err := downloadFile(*url, *output)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Файл успешно загружен.")
	}
}

func downloadFile(url, output string) error {
	// Отправка HTTP-запроса
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Проверка успешного статуса ответа
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("неправильный статус ответа: %s", resp.Status)
	}

	// Создание файла для сохранения данных
	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()

	// Копирование данных из ответа в файл
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}