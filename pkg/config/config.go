package config

import (
	"html/template"
	"log"
)

//AppConfid holds the app config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
