package template

import (
	"time"
	"html/template"
)

var AppHelpers = template.FuncMap{
	"unescaped": unescaped,
	"date": date,
	"time": dateTime,
}

func unescaped(x string) interface{} {
	return template.HTML(x)
}

func date(d time.Time) string {
	return d.Format("2006-01-02")
}

func dateTime(d time.Time) string {
	return d.Format("2006-01-02 15:04:05")
}
