package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Парсинг флагов
	host := flag.String("host", "", "Хост (IP или доменное имя)")
	port := flag.Int("port", 0, "Порт")
	timeout := flag.Duration("timeout", 10*time.Second, "Таймаут на подключение")
	flag.Parse()

	// Проверка обязательных флагов
	if *host == "" || *port == 0 {
		fmt.Println("Необходимо указать хост и порт")
		return
	}

	// Формирование адреса для подключения
	addr := fmt.Sprintf("%s:%d", *host, *port)

	// Установка соединения
	conn, err := net.DialTimeout("tcp", addr, *timeout)
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		return
	}
	defer conn.Close()

	// Канал для обработки сигналов завершения
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	// Запуск горутины для чтения данных из соединения и вывода в STDOUT
	go func() {
		io.Copy(os.Stdout, conn)
		sigCh <- syscall.SIGTERM // отправка сигнала завершения при закрытии соединения
	}()

	// Запуск горутины для чтения данных из STDIN и записи в соединение
	go func() {
		io.Copy(conn, os.Stdin)
		conn.Close() // закрытие соединения при завершении ввода из STDIN
	}()

	// Ожидание сигнала завершения
	<-sigCh
	fmt.Println("\nЗавершение программы.")
}