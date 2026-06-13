package main

type PaginatedResponse struct {
	Data       []User `json:"data"`
	TotalRows  int    `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
}
