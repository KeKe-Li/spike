package controllers

import (
	"net/http"
	"io/ioutil"
	"encoding/json"

	redis_authories "spike/helps/redis"

	"github.com/unrolled/render"
	"github.com/gin-gonic/gin"

)

type AuthorityController struct {
	ApplicationController
}

const (
	AuthorityKey = "ALLIANCE::AUTHORITY"
)

func (ctrl AuthorityController)New(c *gin.Context) {
	ctrl.Render.HTML(c.Writer, http.StatusOK, "authority/index", nil, render.HTMLOptions{Layout: "authority/templates"})
}

func (ctrl AuthorityController)All(c *gin.Context) {
	resp := ctrl.NewResponse()
	allianceRedis := redis_authories.Alliance{Redis:ctrl.RedisPool, Key:AuthorityKey}
	authories := make([]redis_authories.Authority, 0)
	s:=allianceRedis.GetAll()
	err := json.Unmarshal([]byte(s), &authories)
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.ErrorCode = 0
	resp.Data = authories
	c.JSON(http.StatusOK, resp)
}

func (ctrl AuthorityController) Save(c *gin.Context) {
	resp := ctrl.NewResponse()
	data, err := ioutil.ReadAll(ioutil.NopCloser(c.Request.Body))
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = "参数解析错误"
		c.JSON(http.StatusOK, resp)
		return
	}
	defer c.Request.Body.Close()
	authories := make([]redis_authories.Authority, 0)
	err = json.Unmarshal(data, &authories)
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	allianceRedis := redis_authories.Alliance{Redis:ctrl.RedisPool, Key:AuthorityKey}
	err = allianceRedis.Save(string(data))
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.ErrorCode = 0
	resp.Data = authories
	c.JSON(http.StatusOK, resp)
}