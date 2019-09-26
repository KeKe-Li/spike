package models

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type MicroService struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Name    string `json:"name" sql:"size:100;not null;index" ` // not null micro service name and it can not be updated
	Summary string `json:"summary" sql:"size:200;not null"`     //not null micro service description ,it can be updated
	Remark  string `json:"remark" sql:"size:200"`               //remark be used description where gpgkey came from
	GpgKey  []GpgKey
}
