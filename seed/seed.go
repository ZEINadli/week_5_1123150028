package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ZEINadli/week_5_1123150028/config"
	"github.com/ZEINadli/week_5_1123150028/models"
)

func main() {
	godotenv.Load()

	config.InitDatabase()

	products := []models.Product{
		{
			Name: "Es Teh Manis",
			Price: 8000,
			Category: "Minuman",
			Stock: 200,
			Description: "Es teh manis segar",
			ImageURL: "https://picsum.photos/400",
		},
		{
			Name: "Kopi Susu",
			Price: 15000,
			Category: "Minuman",
			Stock: 150,
			Description: "Kopi susu dengan gula aren",
			ImageURL: "https://picsum.photos/401",
		},
		{
			Name: "Jus Alpukat",
			Price: 18000,
			Category: "Minuman",
			Stock: 80,
			Description: "Jus alpukat dengan susu coklat",
			ImageURL: "https://picsum.photos/402",
		},
		{
			Name: "Jus Mangga",
			Price: 17000,
			Category: "Minuman",
			Stock: 90,
			Description: "Jus mangga segar",
			ImageURL: "https://picsum.photos/403",
		},
		{
			Name: "Lemon Tea",
			Price: 12000,
			Category: "Minuman",
			Stock: 120,
			Description: "Teh lemon segar",
			ImageURL: "https://picsum.photos/404",
		},
	}

	for _, p := range products {
		config.DB.Create(&p)
	}

	log.Printf("Seed berhasil: %d produk minuman ditambahkan", len(products))
}