package models

import "time"

type News struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	Created   time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	Title       string     `json:"title" sql:"size:100;not null"`
	Summary     string     `json:"summary" sql:"size:100;not null"`
	Body        string     `json:"source" sql:"type:text;not null"`
	Code        string     `json:"code" sql:"size:21;not null"`
	Author      string     `json:"author" sql:"size:100;"`
	Views       int        `json:"views" sql:"default:0"`
	State       int        `json:"state" sql:"default:1"` //1待审核,2已上线,3已下线
	PublishedAt *time.Time `json:"published_at"`

	ImageUrl      string `json:"image_url" sql:"size:150"`
	ShareImageUrl string `json:"share_image_url" gorm:"-"`
}


func (m *News) TableName() string{
	return "news"
}