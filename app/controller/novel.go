package controller

import (
	"github.com/gin-gonic/gin"
	"go-template/app/model"
	"go-template/app/service"
	"net/http"
	"strconv"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func handleError(c *gin.Context, statusCode int, errorMessage string, err error) {
	c.JSON(statusCode, gin.H{
		"error":   errorMessage,
		"message": err.Error(),
	})
}

func handleSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Message: message,
		Data:    data,
	})
}
func NovelById(c *gin.Context) {
	novelService := service.NovelService{}

	// Extract id from request parameters
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid ID format", err)
		return
	}

	// Fetch novel by id
	novel := novelService.GetNovelById(id)
	if novel == nil {
		handleError(c, http.StatusNotFound, "No novel found with the given ID", nil)
		return
	}

	handleSuccess(c, http.StatusOK, "Novel fetched successfully", novel)
}

func NovelAdd(c *gin.Context) {
	novel := model.Novel{}
	if err := c.ShouldBindJSON(&novel); err != nil {
		handleError(c, http.StatusBadRequest, "Bad Request", err)
		return
	}

	novelService := service.NovelService{}
	id, err := novelService.SetNovel(&novel)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}

	novel.Id = id
	handleSuccess(c, http.StatusCreated, "Novel added successfully", novel)
}

func NovelList(c *gin.Context) {
	novelService := service.NovelService{}
	novelList := novelService.GetNovelList()
	handleSuccess(c, http.StatusOK, "Novel list fetched successfully", novelList)
}

func NovelUpdate(c *gin.Context) {
	novel := model.Novel{}
	if err := c.ShouldBindJSON(&novel); err != nil {
		handleError(c, http.StatusBadRequest, "Bad Request", err)
		return
	}

	novelService := service.NovelService{}
	err := novelService.UpdateNovel(&novel)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}

	handleSuccess(c, http.StatusOK, "Novel updated successfully", novel)
}

func NovelDelete(c *gin.Context) {
	id := c.Query("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		handleError(c, http.StatusBadRequest, "Bad Request", err)
		return
	}

	novelService := service.NovelService{}
	err = novelService.DeleteNovel(intId)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}

	handleSuccess(c, http.StatusOK, "Novel deleted successfully", nil)
}
