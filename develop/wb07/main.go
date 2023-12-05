package main

import (
	"fmt"
	"time"
)

// Функция or принимает один или более done-каналов и возвращает
// single-канал. Если один из составляющих каналов закрывается,
// результатный канал также закрывается
var or func(channels ...<-chan interface{}) <-chan interface{}

func init() {
	// Реализация функции or
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			// Если нет входных каналов, создаем и возвращаем закрытый канал
			c := make(chan interface{})
			close(c)
			return c
		case 1:
			// Если только один входной канал, возвращаем его
			return channels[0]
		}

		// Используем select для слушания всех входных каналов
		orDone := make(chan interface{})
		go func() {
			defer close(orDone)

			select {
			// В случае первого закрытия канала, закрываем orDone
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			// Можно добавить еще случаи для других входных каналов...
			}
		}()

		return orDone
	}
}

func main() {
	// Пример использования функции or
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	// Использование функции or с несколькими done-каналами
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("Done after %v", time.Since(start))
}
