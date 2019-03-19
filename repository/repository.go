package repository

import (
	"bufio"
	"encoding/base64"
	"imagesapi/entity"
	"io/ioutil"
	"log"
	"os"
)

func LoadConvert(c *entity.Image) (err error) {
	r, _ := os.Open("images/rock.jpg")
	if err != nil {
		log.Fatalf("failed to open image to encode: %v", err)
	}

	reader := bufio.NewReader(r)
	content, _ := ioutil.ReadAll(reader)
	encoded := base64.StdEncoding.EncodeToString(content)

	p, _ := os.Open("images/rock64preview.jpg")
	if err != nil {
		log.Fatalf("failed to open image to encode: %v", err)
	}

	reader2 := bufio.NewReader(p)
	content2, _ := ioutil.ReadAll(reader2)
	encoded2 := base64.StdEncoding.EncodeToString(content2)

	r180, _ := os.Open("images/rock180.jpg")
	if err != nil {
		log.Fatalf("failed to open image to encode: %v", err)
	}

	reader3 := bufio.NewReader(r180)
	content3, _ := ioutil.ReadAll(reader3)
	encoded3 := base64.StdEncoding.EncodeToString(content3)

	com := &entity.Image{Name: "rock", ImageStr: encoded, ImageType: "original"}
	err = db.Table("images").Create(com).Error
	if err != nil {
		return
	}
	xom := &entity.Image{Name: "rock64preview", ImageStr: encoded2, ImageType: "preview"}
	err = db.Table("images").Create(xom).Error
	if err != nil {
		return
	}
	dom := &entity.Image{Name: "rock180", ImageStr: encoded3, ImageType: "180"}
	err = db.Table("images").Create(dom).Error
	if err != nil {
		return
	}
	return
}

func GetByFilter(filter *entity.Image) (images []entity.Image, err error) {
	query := db.Table("images")

	if filter.ID != 0 {
		query = query.Where("id=?", filter.ID)
	}

	if filter.ImageType != "" {
		query = query.Where("image_type=?", filter.ImageType)
	}

	if filter.Name != "" {
		query = query.Where("name=?", filter.Name)
	}

	err = query.Find(&images).Error
	return
}
