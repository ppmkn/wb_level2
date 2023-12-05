package utils

import (
	"errors"
	"net/http"
	"strconv"
)

// validateIntParam осуществляет проверку и извлечение целочисленного параметра из запроса
func validateIntParam(r *http.Request, paramName string) (int, error) {
	// Извлекаем значение параметра из формы запроса
	paramValue := r.FormValue(paramName)
	if paramValue == "" {
		// Если параметр отсутствует, возвращаем ошибку
		return 0, errors.New("missing parameter: " + paramName)
	}

	// Преобразуем значение параметра в целое число
	value, err := strconv.Atoi(paramValue)
	if err != nil {
		return 0, errors.New("invalid parameter format: " + paramName)
	}

	return value, nil
}
