package main

import "testing"

func TestUnpack(t *testing.T) {
	// Тестовые случаи с различными входными строками и ожидаемыми результатами.
	tests := []struct {
		input    string
		expected string
		wantErr  bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"qwe\\4\\5", "qwe45", false},
		{"qwe\\45", "qwe44444", false},
		{"qwe\\\\5", "qwe\\\\\\", false},
		{"qwe\\", "", true},
	}

	// Итерация по тестовым случаям.
	for _, test := range tests {
		// Вызов функции Unpack с текущим тестовым случаем.
		result, err := Unpack(test.input)

		// Проверка наличия/отсутствия ошибки в соответствии с ожидаемым результатом.
		if (err != nil) != test.wantErr {
			t.Errorf("Ошибка в тесте %q: ожидалась ошибка %v, получено %v", test.input, test.wantErr, err)
		}

		// Проверка соответствия результата ожидаемому значению.
		if result != test.expected {
			t.Errorf("Ошибка в тесте %q: ожидалось %q, получено %q", test.input, test.expected, result)
		}
	}
}