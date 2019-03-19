package entity

import "time"

type Image struct {
	ID        int       `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	ImageStr  string    `json:"-"`
	ImageType string    `json:"image_type" form:"image_type"`
	CreatedAt time.Time `json:"created_at"`
}
