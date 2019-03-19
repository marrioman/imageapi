package entity

import "time"

type Image struct {
	ID        int       `json:"id" form:"id"`
	Name1     string    `json:"name1" form:"name1"`
	ImageStr  string    `json:"-"`
	ImageType string    `json:"image_type" form:"image_type"`
	CreatedAt time.Time `json:"created_at"`
}
