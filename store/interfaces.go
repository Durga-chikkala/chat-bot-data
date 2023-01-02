package store

import (
	"github.com/Dataservicee/models"
	"github.com/gofiber/fiber/v2"
)

type QueriesEndPoints interface {
	Create(c *fiber.Ctx, data models.QueriesData) (models.QueriesData, error)
	GetByQuestion(c *fiber.Ctx, question string) (models.QueriesData, error)
	Get(c *fiber.Ctx) ([]models.QueriesData, error)
	PatchByQuestion(c *fiber.Ctx, count int64) (models.QueriesData, error)
}
