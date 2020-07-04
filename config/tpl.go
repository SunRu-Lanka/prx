package config

import (
	"html/template"
	"time"
)

var TPL *template.Template
var DbSessionsCleaned time.Time



func init()  {
	TPL = template.Must(template.ParseGlob("templates/*.html"))
	DbSessionsCleaned =time.Now()




}

