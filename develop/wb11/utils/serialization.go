package utils

import (
	"encoding/json"
	"net/http"
)

// JsonResponse отправляет JSON-ответ клиенту с указанным HTTP статусом и данными
func JsonResponse(w http.ResponseWriter, status int, data interface{}) {
	// Устанавливаем заголовок Content-Type для ответа в формате JSON
	w.Header().Set("Content-Type", "application/json")
	// Устанавливаем HTTP статус ответа
	w.WriteHeader(status)

	// Кодируем данные в формат JSON и отправляем их клиенту
	if err := json.NewEncoder(w).Encode(data); err != nil {
		// Если произошла ошибка при кодировании данных, возвращаем HTTP 500
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
