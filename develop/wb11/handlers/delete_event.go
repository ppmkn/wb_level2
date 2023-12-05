package handlers

import (
	"net/http"

	"task11/utils"
)

// DeleteEventHandler обрабатывает запрос на удаление события
func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	// Возвращаем JSON-ответ
	response := map[string]string{"result": "Event deleted successfully"}
	utils.JsonResponse(w, http.StatusOK, response)
}
