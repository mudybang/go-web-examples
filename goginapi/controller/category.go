package controller

import (
	"net/http"

	"github.com/gosimple/slug"

	"github.com/gin-gonic/gin"
	"github.com/mudybang/go-web-examples/goginapi/database"
	"github.com/mudybang/go-web-examples/goginapi/model"
)

func GetAllCategories(context *gin.Context) {
	var categories []model.Category
	err := database.Database.Find(&categories).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": categories})
}

func CreateCategory(context *gin.Context) {
	var input model.Category
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category := model.Category{
		Name: input.Name,
		Slug: slug.Make(input.Name),
	}
	savedCategory, err := category.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedCategory})
}
