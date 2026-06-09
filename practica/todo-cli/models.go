package main

import "time"

type Todo struct {
	ID          int64     `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Kind        int       `json:"kind"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
