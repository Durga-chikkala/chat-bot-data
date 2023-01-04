package main

import (
	"github.com/Dataservicee/handlers"
	"github.com/Dataservicee/store"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	app := fiber.New()
	PORT := 3000

	dsn := "host=localhost user=postgres password=password dbname=queries port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("DB connection failed with Err:", err)
		return
	}

	s := store.New(db)
	h := handlers.New(s)

	app.Post("/db", h.Create)
	app.Get("/db/getAll", h.Get)
	app.Get("/db/:question", h.GetByQuestion)
	app.Patch("/db/:question", h.PatchByQuestion)

	log.Printf("The server is running at port:%v", PORT)
	app.Listen(":3000")
}
