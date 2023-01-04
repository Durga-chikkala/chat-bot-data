package store

import (
	"github.com/Dataservicee/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type database struct {
	db *gorm.DB
}

func New(db *gorm.DB) database {
	return database{db: db}
}

func (d database) Create(c *fiber.Ctx, input models.QueriesData) (models.QueriesData, error) {

	result := d.db.Create(&input)
	if result.Error != nil {
		return models.QueriesData{}, result.Error
	}

	if result.RowsAffected == 0 {
		return models.QueriesData{}, fiber.ErrBadRequest
	}

	return input, nil
}

func (d database) Get(c *fiber.Ctx) ([]models.QueriesData, error) {
	var data []models.QueriesData
	result := d.db.Find(&data)
	if result.Error != nil {
		return []models.QueriesData{}, result.Error
	}

	return data, nil
}

func (d database) GetByQuestion(c *fiber.Ctx, question string) (models.QueriesData, error) {
	var data models.QueriesData

	result := d.db.Find(&data, "question=?", question+"?")
	if result.Error != nil {
		return models.QueriesData{}, result.Error
	}

	if result.RowsAffected == 0 {

		return models.QueriesData{}, fiber.ErrNotFound

	}

	return data, nil
}

func (d database) PatchByQuestion(c *fiber.Ctx, count int64, question string) (models.QueriesData, error) {
	var data models.QueriesData

	result := d.db.Model(&data).Where("question= ?", question).Update("count", count)
	if result.Error != nil {
		return models.QueriesData{}, result.Error
	}

	if result.RowsAffected == 0 {

		return models.QueriesData{}, fiber.ErrNotFound

	}

	return data, nil
}
