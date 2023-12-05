package handlers

import (
	"net/http"

	"task11/utils"
)

// UpdateEventHandler обрабатывает запрос на обновление события
func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	// Возвращаем JSON-ответ
	response := map[string]string{"result": "Event updated successfully"}
	utils.JsonResponse(w, http.StatusOK, response)
}