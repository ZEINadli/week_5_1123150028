package main

import (
	"log"

	"github.com/ZEINadli/week_5_1123150028/config"
	"github.com/ZEINadli/week_5_1123150028/models"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config.InitDatabase()

	products := []models.Product{
		{
			Name:        "Apel Fuji",
			Price:       35000,
			Category:    "Impor",
			Stock:       50,
			Description: "Apel fuji impor dengan rasa manis dan renyah",
			ImageURL:    "https://pasarrakyatbali.com/cdn/shop/products/PicsArt_10-22-04.02.46_1990x.jpg?v=1634889559",
		},
		{
			Name:        "Anggur Merah",
			Price:       45000,
			Category:    "Impor",
			Stock:       40,
			Description: "Anggur merah tanpa biji impor",
			ImageURL:    "https://sevenrose.co.id/cdn/shop/files/Anggur_Red_Globe.jpg?v=1759840360",
		},
		{
			Name:        "Pir Xiang Lie",
			Price:       38000,
			Category:    "Impor",
			Stock:       35,
			Description: "Buah pir impor dengan kandungan air tinggi",
			ImageURL:    "https://i.ibb.co.com/2YStPbQt/PEAR-XIANG-LIE-HD.jpg",
		},
		{
			Name:        "Kiwi Hijau",
			Price:       42000,
			Category:    "Impor",
			Stock:       25,
			Description: "Kiwi impor dengan rasa asam manis",
			ImageURL:    "https://www.zespri.com/content/dam/zespri/id/blog/body-image/eatstoreripen_570x610.jpg",
		},
		{
			Name:        "Jeruk Sunkist",
			Price:       30000,
			Category:    "Impor",
			Stock:       45,
			Description: "Jeruk sunkist impor segar",
			ImageURL:    "https://niagatani.id/uploads/all/r11JLSam7uSiE7yHxnmpox5HJ60xr8Ae3wznXpy3.jpg",
		},
		{
			Name:        "Jeruk Pontianak",
			Price:       18000,
			Category:    "Lokal",
			Stock:       100,
			Description: "Jeruk lokal dengan rasa manis segar",
			ImageURL:    "https://i.ibb.co.com/RtVgtJb/190840-jlcoverwikipedia.jpg",
		},
		{
			Name:        "Mangga Harum Manis",
			Price:       25000,
			Category:    "Lokal",
			Stock:       70,
			Description: "Mangga harum manis asli lokal",
			ImageURL:    "https://i.ibb.co.com/gb42CD39/62a6f23eb2369.jpg",
		},
		{
			Name:        "Pisang Cavendish",
			Price:       22000,
			Category:    "Lokal",
			Stock:       90,
			Description: "Pisang segar dengan tekstur lembut",
			ImageURL:    "https://images.alodokter.com/dk0z4ums3/image/upload/v1741311608/attached_image/pisang-cavendish-ketahui-nutrisi-dan-manfaat-kesehatannya.jpg",
		},
		{
			Name:        "Pepaya California",
			Price:       15000,
			Category:    "Lokal",
			Stock:       60,
			Description: "Pepaya lokal dengan daging manis",
			ImageURL:    "https://i.ibb.co.com/wrP8fHnX/Pepaya-new.jpg",
		},
		{
			Name:        "Salak Pondoh",
			Price:       20000,
			Category:    "Lokal",
			Stock:       80,
			Description: "Salak pondoh dengan rasa manis khas",
			ImageURL:    "https://i.ibb.co.com/qTR3zMC/1669347496-Snake-Fruit-or-Salacca-Salak-Pondoh.png",
		},
	}

	for _, p := range products {
		config.DB.Create(&p)
	}

	log.Printf("Seed berhasil: %d produk minuman ditambahkan", len(products))
}
