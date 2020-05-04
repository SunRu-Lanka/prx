package config

import (
	"html/template"
	"time"
)

var TPL *template.Template
var dbSessionsCleaned time.Time

func init()  {
	TPL = template.Must(template.ParseGlob("templates/*.html"))
	dbSessionsCleaned = time.Now()



}
