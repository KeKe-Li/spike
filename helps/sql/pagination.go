package sql

import (
	"math"
	"net/url"

	"github.com/jinzhu/gorm"
)

type pagination struct {
	Query         *gorm.DB   `json:"-"`
	TotalEntities int32      `json:"total_entities" `
	PerPage       int32      `json:"per_page" `
	Path          string     `json:"path" `
	Page          int32      `json:"page" `
	UrlQuery      url.Values `json:"url_query" `
	TotalPages    int32      `json:"total_pages" `
}

//接受 string int 两种传入值
func (p *pagination) Paginate() *gorm.DB {
	p.Query.Count(&p.TotalEntities)
	if p.TotalEntities == 0 {
		return p.Query
	}
	p.TotalPages = int32(math.Ceil(float64(p.TotalEntities) / float64(p.PerPage)))
	query := p.Query.Offset(int((p.Page - 1) * p.PerPage)).Limit(int(p.PerPage))
	return query

}

func NewPagination(page, per_page int32) *pagination {
	return &pagination{Page: page, PerPage: per_page}
}
