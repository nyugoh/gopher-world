package rest

import "time"

type Post struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Meta        string    `json:"meta"`
	Content     string    `json:"content"`
	ImageCover  string    `json:"image_cover"`
	CreateAt    time.Time `json:"created_at"`
}
