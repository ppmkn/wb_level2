package handlers

import (
	"net/http"

	"task11/utils"
)

// EventsForMonthHandler обрабатывает запрос на получение событий за месяц
func EventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	// Возвращаем JSON-ответ
	response := map[string]string{"result": "Events for the month"}
	utils.JsonResponse(w, http.StatusOK, response)
}
