package domain

type Teacher struct {
	TeacherId int64  `json:"teacher_id"`
	Name      string `json:"Name"`
	Shift     int    `json:"shift"` //1 morning 2 Afternoon
}
