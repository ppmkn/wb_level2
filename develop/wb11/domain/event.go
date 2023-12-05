package domain

// структура доменной области
type Event struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
}