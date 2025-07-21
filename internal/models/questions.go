package models

import "time"

type Question struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	Dimension string    `json:"dimension"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Published bool      `json:"published"`
	Version   int       `json:"version"`
	Locked    bool      `json:"locked"`
}
