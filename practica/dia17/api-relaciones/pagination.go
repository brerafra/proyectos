package main

type PaginatedUserResponse struct {
	Data       []User `json:"data"`
	TotalRows  int    `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
}

type PaginatedPostResponse struct {
	Data       []Post `json:"data"`
	TotalRows  int    `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
}

type PaginatedTaskResponse struct {
	Data       []Task `json:"data"`
	TotalRows  int    `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
}
