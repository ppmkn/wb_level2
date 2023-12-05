package handlers

import (
	"net/http"

	"task11/utils"
)

// EventsForDayHandler обрабатывает запрос на получение событий за день
func EventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	// Возвращаем JSON-ответ
	response := map[string]string{"result": "Events for the day"}
	utils.JsonResponse(w, http.StatusOK, response)
}
