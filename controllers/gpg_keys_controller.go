package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"spike/models"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
)

type GpgKeysController struct {
	ApplicationController
}

func (ctrl GpgKeysController) Index(c *gin.Context) {
	ctrl.Render.HTML(c.Writer, http.StatusOK, "gpg_keys/index", gin.H{
		"title": "index",
	}, render.HTMLOptions{Layout: "template_layout"})

}

func (ctrl GpgKeysController) Show(c *gin.Context) {
	var gpgKey []models.GpgKey
	msid := c.Param("id")

	err := ctrl.DB.Where("micro_service_id = ?", msid).Find(&gpgKey).Error
	if err != nil {
		fmt.Println(err)
	}
	ctrl.Render.HTML(c.Writer, http.StatusOK, "gpg_keys/show", gin.H{
		"title":  "show",
		"gpgKey": gpgKey,
		"msid":   msid,
	}, render.HTMLOptions{Layout: "template_layout"})

}

func (ctrl GpgKeysController) Detail(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 0)
	gpgKey := models.GpgKey{ID:uint(id)}
	err := ctrl.DB.First(&gpgKey).Error
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "gpg_keys/detail", gin.H{
			"title":  "detail",
			"gpgKey": gpgKey,
			"msid":   id,
		}, render.HTMLOptions{Layout: "template_layout"})
	}

}

func (ctrl GpgKeysController) New(c *gin.Context) {
	var microService models.MicroService
	msid := c.Param("id")
	err := ctrl.DB.Where("id = ?", msid).Find(&microService).Error
	if err != nil {
		fmt.Println(err)
	}
	ctrl.Render.HTML(c.Writer, http.StatusOK, "gpg_keys/new", gin.H{
		"title":  "new",
		"msid":   msid,
		"msName": microService.Name,
	}, render.HTMLOptions{Layout: "template_layout"})
}

func (ctrl GpgKeysController) Create(c *gin.Context) {


}


func (ctrl GpgKeysController) Edit(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
		return
	}

	gpgKey := models.GpgKey{ID: uint(id)}
	err = ctrl.DB.First(&gpgKey).Error
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "gpg_keys/edit", gin.H{
			"title":  "edit",
			"gpgKey": gpgKey,
		})
	}
}

func (ctrl GpgKeysController) Update(c *gin.Context) {
	resp := ctrl.NewResponse()

	armor, err := strconv.ParseBool(c.PostForm("armor"))
	if err !=nil{
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
		return
	}
	comment := c.PostForm("comment")
	email := c.PostForm("email")
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil{
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
		return
	}
	microServiceId, err := strconv.ParseUint(c.PostForm("microServiceId"), 10, 0)
	name := c.PostForm("name")
	pubRing := c.PostForm("pubRing")
	secRing := c.PostForm("secRing")
	remoteKey := c.PostForm("remoteKey")

	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
		return
	}

	updateColumns := make(map[string]interface{})
	updateColumns["armor"] = armor
	updateColumns["comment"] = comment
	updateColumns["email"] = email
	updateColumns["micro_service_id"] = microServiceId
	updateColumns["name"] = name
	updateColumns["pub_ring"] = pubRing
	updateColumns["sec_ring"] = secRing
	updateColumns["remoteKey"] = remoteKey

	gpgKey := &models.GpgKey{
		ID: uint(id),
	}

	err = ctrl.DB.Model(gpgKey).UpdateColumns(updateColumns).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}
}

func (ctrl GpgKeysController) Delete(c *gin.Context) {
	resp := ctrl.NewResponse()

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
		return
	}

	gpgKey := &models.GpgKey{}

	err = ctrl.DB.Model(gpgKey).Delete(gpgKey, id).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}
}

