package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
	var err error

	for {
		db, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			"127.0.0.1",
			"5432",
			"postgres",
			"postgres",
			"postgres",
			"disable",
		))

		if err != nil {
			log.Println(err)
			time.Sleep(time.Second * 5)
			continue
		}
		db.SingularTable(true)
		db.LogMode(true)
		migration()
		break
	}
}

func migration() {
	driver, err := postgres.WithInstance(db.DB(), &postgres.Config{})
	if err != nil {
		log.Printf("failure migration driver %s", err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		log.Printf("failure migration file %s", err)
		return
	}

	if err := m.Up(); err != nil {
		log.Printf("failure migration up %s", err)
		return
	}
}
