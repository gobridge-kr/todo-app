package model

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Order     int64  `json:"order"`
	URL       string `json:"url"`
}
