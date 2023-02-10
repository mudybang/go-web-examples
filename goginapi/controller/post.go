package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mudybang/go-web-examples/goginapi/helper"
	"github.com/mudybang/go-web-examples/goginapi/model"
)

func GetAllPosts(context *gin.Context) {
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user.Posts})
}

func CreatePost(context *gin.Context) {
	var input model.Post
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = user.ID

	savedPost, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedPost})
}
