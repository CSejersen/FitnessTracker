package models

type Program struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Name    string `json:"name"`
	Split   string `json:"split"`
	PerWeek int    `json:"per_week"`
}
