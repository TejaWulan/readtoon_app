package main

import (
	"log"
	"readtoon_app/config"
	"readtoon_app/database"
	"readtoon_app/models"
	"readtoon_app/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Koneksi DB
	database.InitDB(cfg)

	// Auto migrate
	database.DB.AutoMigrate(
		&models.AccountProfile{},
		&models.Canvas{},
		&models.CanvasEpisodeSummary{},
		&models.CanvasMusicVideoWork{},
		&models.CanvasWorkLibrary{},
		&models.Coin{},
		&models.CoinPurchaseTransaction{},
		&models.Comic{},
		&models.ComicEpisodeSummary{},
		&models.Episode{},
		&models.Library{},
		&models.LibraryItem{},
		&models.MusicVideo{},
		&models.MusicVideoFavouriteItem{},
		&models.MusicVideoFavouriteLibrary{},
		&models.MusicVideoSummary{},
		&models.MusicVideoWorkLibrary{},
		&models.PaymentMethod{},
	)

	log.Println("Database siap!")

	// Jalankan server
	e := echo.New()
	routes.InitRoutes(e)
	e.Logger.Fatal(e.Start(cfg.SERVER_HOST))
}
