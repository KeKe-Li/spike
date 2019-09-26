package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"spike/helps/sql"
	"spike/models"

	"github.com/gin-gonic/gin"
)

type NewsController struct {
	ApplicationController
}

func (ctrl NewsController) List(c *gin.Context) {
	resp := ctrl.NewResponse()

	page, err := strconv.ParseInt(c.Query("page"), 10, 0)
	if err != nil {
		page = 1
	}

	perPage, err := strconv.ParseInt(c.Query("perPage"), 10, 0)
	if err != nil {
		perPage = 20
	}

	orderKey := c.Query("order_key")
	if orderKey == "" {
		orderKey = "created_at"
	}

	orderValue := c.Query("order_value")
	if orderValue == "" {
		orderValue = "desc"
	}

	var state int32
	switch c.Query("state") {
	case "publish":
		state = int32(models.Pushlish)
	case "offline":
		state = int32(models.OffLine)
	case "pendding":
		state = int32(models.Pendding)
	case "all":
		state = 0
	}

	req := struct {
		Page       int64  `json:"page"`
		PerPage    int64  `json:"per_page"`
		OrderKey   string `json:"order_key"`
		OrderValue string `json:"order_value"`
		State      int32  `json:"state"`
	}{
		Page:       page,
		PerPage:    perPage,
		OrderValue: orderValue,
		OrderKey:   orderKey,
		State:      state,
	}

	conditions := make(map[string]interface{})
	if req.State > 0 {
		conditions["state"] = req.State
	}
	query := ctrl.DB.Model(&models.News{})
	if len(conditions) > 0 {
		query = query.Where(conditions)
	}

	pagination := sql.NewPagination(int32(req.Page), int32(req.PerPage))
	pagination.Query = query.Order(fmt.Sprintf("%s %s", req.OrderKey, req.OrderValue))

	n := make([]models.News, 0)
	err = pagination.Paginate().Find(&n).Error
	if err != nil {
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	result := struct {
		Data     interface{} `json:"data"`
		PageInfo interface{} `json:"page_info"`
	}{
		Data:     n,
		PageInfo: pagination,
	}

	resp.ErrorCode = 0
	resp.Data = result
	c.JSON(http.StatusOK, resp)

}

func (ctrl NewsController) Show(c *gin.Context) {

}

func (ctrl NewsController) Publish(c *gin.Context) {
	resp := ctrl.NewResponse()

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	req := &struct{ ID uint64 }{ID: id}
	n := &models.News{}
	err = ctrl.DB.Model(n).First(n, req.ID).Error
	if err != nil {
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	if n.State == int(models.Pendding) {
		resp.ErrorMessage = "news state is pendding"
		c.JSON(http.StatusOK, resp)
		return
	}

	err = ctrl.DB.Model(&models.News{}).Where("id=?", req.ID).Update("state", models.Pushlish).Error
	if err != nil {
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.ErrorCode = 0
	c.JSON(http.StatusOK, resp)
}

func (ctrl NewsController) Offline(c *gin.Context) {
	resp := ctrl.NewResponse()

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	req := &struct{ ID uint64 }{ID: id}

	n := &models.News{}

	err = ctrl.DB.Model(n).First(n, req.ID).Error
	if err != nil {
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	if n.State == int(models.Pushlish) {
		resp.ErrorMessage = "news state is publish"
		c.JSON(http.StatusOK, resp)
		return
	}

	err = ctrl.DB.Model(&models.News{}).Where("id=? AND state=?", req.ID, models.Pushlish).
		Update("state", models.OffLine).Error
	if err != nil {
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.ErrorCode = 0
	c.JSON(http.StatusOK, resp)
}

func (ctrl NewsController) Create(c *gin.Context) {
	resp := ctrl.NewResponse()

	title := c.PostForm("title")
	summary := c.PostForm("summary")
	body := c.PostForm("body")
	code := c.PostForm("code")
	author := c.PostForm("author")
	imageUrl := c.PostForm("imageUrl")

	news := models.News{
		Title:    title,
		Summary:  summary,
		Body:     body,
		Code:     code,
		Author:   author,
		ImageUrl: imageUrl,
	}

	err := ctrl.DB.Create(&news).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.ErrorCode = 0
	c.JSON(http.StatusOK, resp)

}
