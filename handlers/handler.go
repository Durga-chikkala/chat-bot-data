package handlers

import (
	"log"

	"github.com/Dataservicee/models"
	"github.com/Dataservicee/store"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Queries store.QueriesEndPoints
}

func New(Q store.QueriesEndPoints) Handler {
	return Handler{Queries: Q}
}

func (h Handler) Create(c *fiber.Ctx) error {
	var input models.QueriesData

	err := c.BodyParser(&input)
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = validateInput(input)
	if err != nil {
		c.Status(400)
		return err
	}

	resp, err := h.Queries.Create(c, input)
	c.JSON(resp)

	return err
}

func (h Handler) Get(c *fiber.Ctx) error {
	resp, err := h.Queries.Get(c)
	c.JSON(resp)
	return err
}

func (h Handler) GetByQuestion(c *fiber.Ctx) error {
	question := c.Params("question")
	log.Print("question:", question)

	if question == "" {
		return fiber.ErrBadRequest
	}

	resp, err := h.Queries.GetByQuestion(c, question)
	c.JSON(resp)

	return err
}

func (h Handler) PatchByQuestion(c *fiber.Ctx) error {
	question := c.Params("question")
	log.Print("question:", question)

	if question == "" {
		return fiber.ErrBadRequest
	}

	var input models.QueriesData
	err := c.BodyParser(&input)
	if err != nil {
		return fiber.ErrBadRequest
	}

	resp, err := h.Queries.PatchByQuestion(c, input.Count)
	c.JSON(resp)
	return err
}

func validateInput(input models.QueriesData) error {
	if input.Question == "" {
		return fiber.EmptyFieldError{Key: "title"}
	}

	if input.Solution == "" {
		return fiber.EmptyFieldError{Key: "Solution"}
	}

	return nil
}
