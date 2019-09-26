package controllers

import (
	"encoding/json"
	"fmt"

	"spike/config"

	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/unrolled/render"
)

type ApplicationController struct {
	RedisPool *redis.Pool
	DB        *gorm.DB
	Render    *render.Render
	Config    *config.ApplicationConfig
}

func (a ApplicationController) Log(v ...interface{}) {
	fmt.Println(v)
}

type ApiResponse struct {
	ErrorCode    int         `json:"error_code"`
	ErrorMessage string      `json:"error_message"`
	Data         interface{} `json:"response,omitempty"`
}

func (act ApplicationController) NewResponse() *ApiResponse {
	return &ApiResponse{}
}

func (a *ApiResponse) ToJSON() []byte {
	result, _ := json.Marshal(a)
	return result
}
