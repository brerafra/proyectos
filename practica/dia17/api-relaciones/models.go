package main

import "time"

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
	IsAdmin  bool   `json:"is_admin"`
	Tasks    []Task `json:"tasks,omitempty"`
}

type Post struct {
	ID      int64  `json:"id"`
	UserId  int64  `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Task struct {
	ID        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Kind      int       `json:"kind"`
	CreatedAt time.Time `json:"created_at"`
	//UpdatedAt time.Time `json:"updated_at"`
}
