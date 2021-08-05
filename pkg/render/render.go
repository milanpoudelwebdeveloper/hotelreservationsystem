package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to the browser", err)
	}

}

//CreateTemplateCache creates a template as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
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
