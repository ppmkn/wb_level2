package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	// Получаем текущее локальное время
	localTime := time.Now()
	fmt.Println("Текущее локальное время:", localTime)

	// Получаем точное время с использованием NTP
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		// Если произошла ошибка, выводим ее в STDERR и выходим с ненулевым кодом
		fmt.Fprintln(os.Stderr, "Ошибка получения времени с NTP:", err)
		os.Exit(1)
	}

	fmt.Println("Точное время с использованием NTP:", ntpTime)
}
