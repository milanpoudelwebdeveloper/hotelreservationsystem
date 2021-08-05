package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error here is:", err)
		return
	}

}

//CreateTemplateCache creates a template as a map
func CreateTemplateCache(w http.ResponseWriter) (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	fmt.Println("the pages are:", pages)
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		fmt.Println("orginal page is", page)
		name := filepath.Base(page)
		fmt.Println("check the filenames:", name)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		fmt.Println("template is", ts)
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
	fmt.Println("let's see the cache", myCache)
	return myCache, nil
}
