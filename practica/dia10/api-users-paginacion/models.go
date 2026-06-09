package main

type User struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Pin    string `json:"pin"`
	Status int    `json:"status"`
}
