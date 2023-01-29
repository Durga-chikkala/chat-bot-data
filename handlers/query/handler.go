package query

import (
	"fmt"
	"github.com/Dataservicee/errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"

	"github.com/Dataservicee/models"
	"github.com/Dataservicee/store"
)

type Handler struct {
	Queries store.QueriesEndPoints
}

func New(Q store.QueriesEndPoints) Handler {
	return Handler{Queries: Q}
}

func (h Handler) Create(c *gin.Context) {
	var input models.QueryInfo

	err := c.Bind(&input)
	if err != nil {
		c.JSON(400, errors.ErrorResponse{Code: "BAD REQUEST", Reason: "Invalid Body"})
		return
	}

	err = validateInput(&input)
	if err != nil {
		c.JSON(400, err)
		return
	}

	resp, err := h.Queries.Create(c, input)
	if err != nil {
		err := err.(errors.ErrorResponse)
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h Handler) Get(c *gin.Context) {
	resp, _ := h.Queries.Get(c)
	c.JSON(http.StatusOK, resp)
}

func (h Handler) GetByQuestion(c *gin.Context) {
	question := strings.TrimSpace(c.Param("question"))
	log.Print("question:", question)

	if question == "" {
		c.JSON(400, errors.ErrorResponse{Code: "BAD REQUEST", Reason: fmt.Sprintf("Missing parameter %v", "question")})
		return
	}

	question = refactorQuestion(question)

	resp, _ := h.Queries.GetByQuestion(c, question)
	c.JSON(http.StatusOK, resp)
}

func (h Handler) PatchByQuestion(c *gin.Context) {
	question := c.Param("question")
	log.Print("question:", question)

	if question == "" {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{Code: "BAD REQUEST", Reason: fmt.Sprintf("Missing parameter %v", "question")})
		return
	}

	question = refactorQuestion(question)

	var input models.QueryInfo
	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{Code: "BAD REQUEST", Reason: fmt.Sprintf("Missing parameter %v", "question")})
		return
	}

	resp, err := h.Queries.PatchByQuestion(c, input.Count, question)
	c.JSON(http.StatusOK, resp)
	return
}

func validateInput(input *models.QueryInfo) error {
	if input.Question == "" {
		return errors.ErrorResponse{Code: "BAD REQUEST", Reason: "Missing Field Question"}
	}

	if input.Solution == "" {
		return errors.ErrorResponse{Code: "BAD REQUEST", Reason: "Missing Field Solution"}
	}

	input.Question = refactorQuestion(input.Question)

	return nil
}

func refactorQuestion(question string) string {

	question = strings.TrimSpace(question)
	return strings.ReplaceAll(question, " ", "-")
}
