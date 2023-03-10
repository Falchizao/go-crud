package models

type Todo struct {
	ID      int64  `json:"id"`
	Subject string `json:"subject"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}
