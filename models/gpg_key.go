package models

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type GpgKey struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	MicroServiceID uint   `json:"micro_service_id" sql:"not null"` //not null micro service unique identifier
	Name           string `json:"name" sql:"size:20;not null"`     //not null gpgkey name
	Comment        string `json:"comment" sql:"size:30;not null"`  // not null gpgkey description
	Email          string `json:"email" sql:"size:50;not null"`    // not null it be used contact developer
	Armor          bool   `json:"armor" sql:"default:1"`           //default value is ture , pubRing will be Binary when it be false
	EndPoint       string `json:"end_point" sql:"not null"`
	RemoteKey      string `json:"remote_name" sql:"size:150;not null"` // not null

	Config  string `json:"config" sql:"not null"`          //加一个字段，config的内容
	Version string `json:"version" sql:"size:30;not null"` // not null micro service version,it can be updated

	PubRing string `json:"pubring" sql:"not null"` //not null and readonly public key and it can be downloaded
	SecRing string `json:"secring" sql:"not null"` // not null and readonly private key
}
