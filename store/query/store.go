package query

import (
	"github.com/Dataservicee/errors"
	"github.com/Dataservicee/models"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

type database struct {
	db *gorm.DB
}

func New(db *gorm.DB) database {
	return database{db: db}
}

func (d database) Create(c *gin.Context, input models.QueryInfo) (models.QueryInfo, error) {

	result := d.db.Create(&input)
	if result.Error != nil {
		return models.QueryInfo{}, errors.ErrorResponse{StatusCode: http.StatusInternalServerError, Code: "INTERNAL SERVER ERROR", Reason: "DB ERROR"}
	}

	if result.RowsAffected == 0 {
		return models.QueryInfo{}, errors.ErrorResponse{StatusCode: http.StatusBadRequest, Code: "BAD REQUEST", Reason: "Unable to Insert Data"}
	}

	return input, nil
}

func (d database) Get(c *gin.Context) ([]models.QueryInfo, error) {
	var data []models.QueryInfo
	result := d.db.Find(&data)
	if result.Error != nil {
		return []models.QueryInfo{}, result.Error
	}

	return data, nil
}

func (d database) GetByQuestion(c *gin.Context, question string) (models.QueryInfo, error) {
	var data models.QueryInfo

	result := d.db.Find(&data, "question=?", question+"?")
	if result.Error != nil {
		return models.QueryInfo{}, result.Error
	}

	if result.RowsAffected == 0 {

		return models.QueryInfo{}, fiber.ErrNotFound

	}

	return data, nil
}

func (d database) PatchByQuestion(c *gin.Context, count int64, question string) (models.QueryInfo, error) {
	var data models.QueryInfo

	result := d.db.Model(&data).Where("question=?", question).Update("count", count)
	if result.Error != nil {
		return models.QueryInfo{}, result.Error
	}

	if result.RowsAffected == 0 {

		return models.QueryInfo{}, fiber.ErrNotFound

	}

	return data, nil
}
