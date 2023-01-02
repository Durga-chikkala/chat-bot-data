package store

import (
	"github.com/Dataservicee/models"
	"github.com/gofiber/fiber/v2"
)

type database struct {
}

func New() database {
	return database{}
}

func (d database) Create(c *fiber.Ctx, input models.QueriesData) (models.QueriesData, error) {

	return input, nil
}

func (d database) Get(c *fiber.Ctx) ([]models.QueriesData, error) {

	return []models.QueriesData{{Id: "2", Count: 4}}, nil
}

func (d database) GetByQuestion(c *fiber.Ctx, question string) (models.QueriesData, error) {

	return models.QueriesData{Id: "2", Count: 4}, nil
}

func (d database) PatchByQuestion(c *fiber.Ctx, count int64) (models.QueriesData, error) {

	return models.QueriesData{Id: "2", Count: count}, nil
}
