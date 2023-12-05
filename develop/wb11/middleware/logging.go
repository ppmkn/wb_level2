package middleware

import (
	"log"
	"net/http"
)

// LoggingMiddleware принимает next http.Handler и возвращает новый обработчик,
// который выполняет логирование перед передачей запроса следующему обработчику
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Логируем информацию о запросе, включая метод и путь
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		// Передаем запрос следующему обработчику
		next.ServeHTTP(w, r)
	})
}