package controllers

import (
	"spike/models"

	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	ApplicationController
}

func (ctrl UserController) Index(c *gin.Context) {
	resp := ctrl.NewResponse()
	var users []models.User

	err := ctrl.DB.Model(&models.User{}).Find(&users).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.ErrorCode = 0
	c.JSON(http.StatusOK, resp)
	return
}

func (ctrl UserController) Create(c *gin.Context) {
	resp := ctrl.NewResponse()

	name := c.PostForm("name")
	email := c.PostForm("email")
	encryptedPassword := c.PostForm("encryptedPassword")

	user := models.User{
		Name:              name,
		Email:             email,
		EncryptedPassword: encryptedPassword,
	}

	err := ctrl.DB.Model(&user).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.ErrorCode = 0
	c.JSON(http.StatusOK, resp)
}
