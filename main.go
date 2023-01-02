package main

import (
	"github.com/Dataservicee/handlers"
	"github.com/Dataservicee/store"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	PORT := 3000

	s := store.New()
	h := handlers.New(s)

	app.Post("/db", h.Create)
	app.Get("/db", h.Get)
	app.Get("/db/:question", h.GetByQuestion)
	app.Patch("/db/:question", h.PatchByQuestion)

	log.Printf("The server is running at port:%v", PORT)
	app.Listen(":3000")
}
