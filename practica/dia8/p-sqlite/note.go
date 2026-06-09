package main

import "time"

type Note struct {
	ID          int64     `json:"id,omitempty"`
	Title       string    `json:"title"`
	Despcrition string    `json:"description"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
