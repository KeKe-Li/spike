package main

import (
	"html/template"
	"net/http"

	"spike/config"
	"spike/controllers"
	et "spike/helps/template"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/unrolled/render"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/contrib/static"
)

var (
	appConfig     *config.ApplicationConfig
	router        *gin.Engine
	genericRender *render.Render
	sqlite        *gorm.DB
)

func main() {
	gin.SetMode(appConfig.ENV)
	router = gin.Default()
	router.Use(static.Serve("/static", static.LocalFile("public", true)))

	appCtrl := controllers.ApplicationController{
		Config :appConfig,
	}

	genericRender = render.New(render.Options{
		Directory:  "templates", // Specify what path to load the templates from.
		Layout:     "application_layout", // Specify a layout template. Layouts can call {{ yield }} to render the current template or {{ partial "css" }} to render a partial from the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		Funcs:     []template.FuncMap{et.AppHelpers}, // Specify helper function maps for templates to access.
		Delims:     render.Delims{"{{", "}}"}, // Sets delimiters to the specified strings.
		Charset:    "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: false, // Output human readable JSON.
		IndentXML:  false, // Output human readable XML.
		//PrefixJSON: []byte(")]}',\n"), // Prefixes JSON responses with the given bytes.
		PrefixXML: []byte("<?xml version='1.0' encoding='UTF-8'?>"), // Prefixes XML responses with the given bytes.
		//HTMLContentType: "application/xhtml+xml", // Output XHTML content type instead of default "text/html".
		IsDevelopment:             appConfig.ENV == gin.DebugMode, // Render will now recompile the templates on every HTML response.
		UnEscapeHTML:              true, // Replace ensure '&<>' are output correctly (JSON only).
		StreamingJSON:             true, // Streams the JSON response via json.Encoder.
		RequirePartials:           true, // Return an error if a template is missing a partial used in a layout.
		DisableHTTPErrorRendering: true, // Disables automatic rendering of http.StatusInternalServerError when an error occurs.
	})

	appCtrl.Render = genericRender
	appCtrl.DB = sqlite

	router.GET("", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "")
	})

	spikeRouter := router.Group("/")
	InitRoutes(spikeRouter, appCtrl)
	//权限
	if gin.Mode() == "release" {
		basicAuth := spikeRouter.Group("/authority", gin.BasicAuth(gin.Accounts{
			"spike":    "spike",
		}))
		{
			authCtrl := controllers.AuthorityController{appCtrl}
			basicAuth.GET("/new", authCtrl.New)
			basicAuth.GET("/All", authCtrl.All)
			basicAuth.POST("/save", authCtrl.Save)
		}
	}

	endless.ListenAndServe(":" + appConfig.Port, router)

}