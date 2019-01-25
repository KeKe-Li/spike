package controllers

import (
	"net/http"
	"strconv"

	"spike/models"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
)

type MicroServicesController struct {
	ApplicationController
}

func (ctrl MicroServicesController) Index(c *gin.Context) {
	var microServices []models.MicroService

	err := ctrl.DB.Find(&microServices).Error
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "micro_services/index", gin.H{
			"title":         "index",
			"microServices": microServices,
		}, render.HTMLOptions{Layout: "template_layout"})
	}
}

func (ctrl MicroServicesController) Show(c *gin.Context) {
	var microService models.MicroService
	err := ctrl.DB.Find(&microService).Error
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "micro_services/show", gin.H{
			"title":        "show",
			"microService": microService,
		}, render.HTMLOptions{Layout: "template_layout"})
	}

}

func (ctrl MicroServicesController) New(c *gin.Context) {
	ctrl.Render.HTML(c.Writer, http.StatusOK, "micro_services/new", gin.H{
		"title": "new",
	}, render.HTMLOptions{Layout: "template_layout"})
}

func (ctrl MicroServicesController) Create(c *gin.Context) {
	resp := ctrl.NewResponse()
	name := c.PostForm("name")
	remark := c.PostForm("remark")
	summary := c.PostForm("summary")

	micro_service := models.MicroService{
		Name:    name,
		Remark:  remark,
		Summary: summary,
	}
	err := ctrl.DB.Create(&micro_service).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}

}

func (ctrl MicroServicesController) Edit(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
		return
	}

	microService := models.MicroService{ID: uint(id)}
	err = ctrl.DB.First(&microService).Error
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "micro_services/edit", gin.H{
			"title":        "edit",
			"microService": microService,
		}, render.HTMLOptions{Layout: "template_layout"})
	}
}

func (ctrl MicroServicesController) Update(c *gin.Context) {
	resp := ctrl.NewResponse()

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
		return
	}
	remark := c.PostForm("remark")
	summary := c.PostForm("summary")
	version := c.PostForm("version")

	updateColumns := make(map[string]interface{})
	updateColumns["remark"] = remark
	updateColumns["summary"] = summary
	updateColumns["version"] = version

	microService := &models.MicroService{
		ID: uint(id),
	}
	err = ctrl.DB.Model(microService).UpdateColumns(updateColumns).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}
}

func (ctrl MicroServicesController) Delete(c *gin.Context) {
	resp := ctrl.NewResponse()

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
		return
	}

	microService := &models.MicroService{}
	err = ctrl.DB.Model(microService).Delete(microService, id).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}
}
