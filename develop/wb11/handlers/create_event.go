package handlers

import (
	"net/http"

	"task11/utils"
)

// CreateEventHandler обрабатывает запрос на создание события
func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	// Возвращаем JSON-ответ
	response := map[string]string{"result": "Event created successfully"}
	utils.JsonResponse(w, http.StatusOK, response)
}