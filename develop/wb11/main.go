package main

import (
	"log"
	"net/http"

	"task11/middleware"
	"task11/handlers"
)

func main() {
	// Создаем новый маршрутизатор для обработки HTTP-запросов
	router := http.NewServeMux()

	// Регистрируем обработчики для каждого API-метода
	router.HandleFunc("/create_event", handlers.CreateEventHandler)
	router.HandleFunc("/update_event", handlers.UpdateEventHandler)
	router.HandleFunc("/delete_event", handlers.DeleteEventHandler)
	router.HandleFunc("/events_for_day", handlers.EventsForDayHandler)
	router.HandleFunc("/events_for_week", handlers.EventsForWeekHandler)
	router.HandleFunc("/events_for_month", handlers.EventsForMonthHandler)

	// Используем middleware.LoggingMiddleware для логирования запросов
	// и передаем ему наш маршрутизатор
	http.Handle("/", middleware.LoggingMiddleware(router))

	log.Println("Server is running on :8080")

	// Запускаем HTTP-сервер на порту 8080 и обрабатываем возможные ошибки
	log.Fatal(http.ListenAndServe(":8080", nil))
}
