package main

import (
	"imagesapi/middleware"
	"log"
	"time"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

type Image struct {
	ID        int
	Name1     string
	ImageStr  string
	ImageType string
	CreatedAt time.Time
}

func main() {

	ImageConv()

	HandleRequest()

}

func ImageConv() {

	rev, err := imaging.Open("images/rock.jpg")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	rev = imaging.Rotate180(rev)

	prev := imaging.Resize(rev, 64, 64, imaging.Box)

	err = imaging.Save(prev, "images/rock64preview.jpg")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

	err = imaging.Save(rev, "images/rock180.jpg")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

}

func HandleRequest() {
	r := gin.Default()
	r.GET("/image", middleware.GetByFilterIMG)
	r.POST("/load", middleware.LoadConvertIB)
	r.Run(":8081")
}
