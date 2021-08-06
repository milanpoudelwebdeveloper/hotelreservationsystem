package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

//AppConfid holds the app config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
