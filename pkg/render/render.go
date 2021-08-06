package render

import (
	"bytes"
	"fmt"
	"hotelsystem/pkg/config"
	"hotelsystem/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

//NewTemplate sets the config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a
}
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	//get the template from app config
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to the browser", err)
	}

}

//CreateTemplateCache creates a template as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	//fmt.Println("the pages are:", pages)
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		//fmt.Println("orginal page is", page)
		name := filepath.Base(page)
		//	fmt.Println("check the filenames:", name)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		//	fmt.Println("template is", ts)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts

	}
	//	fmt.Println("let's see the cache", myCache)
	return myCache, nil
}
