package models

type Workout struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id"`
	Name      string     `json:"name"`
	Exercises []Exercise `json:"exercises"`
}
