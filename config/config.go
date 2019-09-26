package config

import (
	"spike/helps/configurations"

	"github.com/jinzhu/gorm"
	"github.com/unrolled/render"
)

type ApplicationConfig struct {
	ENV    string
	Port   string
	DB     *gorm.DB
	Render *render.Render
	Mysql  *configurations.Mysql
	Sqlite string
}
