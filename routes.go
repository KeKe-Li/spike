package main

import (
	"spike/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoute(engine *gin.RouterGroup, appCtrl controllers.ApplicationController) {
	router := engine.Group("")
	{
		//gpgKeys
		gpgKeysCtrl := controllers.GpgKeysController{appCtrl}
		router.GET("/gpg-keys", gpgKeysCtrl.Index)
		router.GET("/gpg-keys-show/:id", gpgKeysCtrl.Show)
		router.GET("/gpg-keys-detail/:id", gpgKeysCtrl.Detail)
		router.GET("/gpg-keys/new/:id", gpgKeysCtrl.New)
		router.GET("/gpg-keys-edit/:id", gpgKeysCtrl.Edit)
		router.POST("/gpg-keys/:id", gpgKeysCtrl.Create)
		//router.POST("/gpg-download/:id",gpgKeysCtrl.Download)
		//router.GET("/gpg-download/:id",gpgKeysCtrl.Download)
		router.PUT("/gpg-keys/:id", gpgKeysCtrl.Update)
		router.DELETE("/gpg-keys/:id", gpgKeysCtrl.Delete)

		//MicroService
		microServicesCtrl := controllers.MicroServicesController{appCtrl}
		router.GET("/micro-services", microServicesCtrl.Index)
		router.GET("/micro-services-show/:id", microServicesCtrl.Show)
		router.GET("/micro-services/new", microServicesCtrl.New)
		router.GET("/micro-services-edit/:id", microServicesCtrl.Edit)
		router.POST("/micro-services", microServicesCtrl.Create)
		router.PUT("/micro-services/:id", microServicesCtrl.Update)
		router.DELETE("/micro-services/:id", microServicesCtrl.Delete)

		userCtrl := controllers.UserController{appCtrl}
		router.GET("/user", userCtrl.Index)
		router.POST("/user", userCtrl.Create)

		newsCtrl := controllers.NewsController{appCtrl}
		router.GET("/news", newsCtrl.List)
		//router.GET("/news-show/:id", newsCtrl.Show)
		router.PUT("/news/publish/:id", newsCtrl.Publish)
		//router.PUT("/news/offline/:id", newsCtrl.OffLine)
		//router.PUT("/news/update/:code", newsCtrl.Update)
		router.POST("/news", newsCtrl.Create)
		//router.POST("/news/save/:code", newsCtrl.Save)
	}

}
