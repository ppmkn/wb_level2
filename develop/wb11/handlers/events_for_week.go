package handlers

import (
	"net/http"

	"task11/utils"
)

// EventsForWeekHandler обрабатывает запрос на получение событий за неделю
func EventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	// Возвращаем JSON-ответ
	response := map[string]string{"result": "Events for the week"}
	utils.JsonResponse(w, http.StatusOK, response)
}