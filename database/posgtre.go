package database

import (
	"fmt"
	"log"
	"readtoon_app/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *config.Configuration) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.PGSQL_HOST,
		cfg.PGSQL_USER,
		cfg.PGSQL_PASSWORD,
		cfg.PGSQL_DATABASE,
		cfg.PGSQL_PORT,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal connect ke database: ", err)
	}
	fmt.Println("Berhasil konek ke PostgreSQL!")
}
