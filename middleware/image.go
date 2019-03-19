package middleware

import (
	"imagesapi/entity"
	"imagesapi/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadConvertIB(c *gin.Context) {
	var image entity.Image

	if err := c.ShouldBindJSON(&image); err != nil {
		println(err.Error())
	}

	if err := repository.LoadConvert(&image); err != nil {
		println(err.Error())
	}
	c.JSON(http.StatusOK, image)
}

func GetByFilterIMG(c *gin.Context) {
	var filter entity.Image

	if err := c.ShouldBindQuery(&filter); err != nil {
		println(err.Error())
	}

	image, err := repository.GetByFilter(&filter)
	log.Println(filter)
	if err != nil {
		println(err.Error())
	}

	c.JSON(http.StatusOK, image)
}
