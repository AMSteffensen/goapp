package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// RenderTemplateTest renders a template using html template
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	if err != nil {
		log.Println("error parsing template:", err)
		return
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("error executing template:", err)
	}
}

var tc = make(map[string]*template.Template)

// RenderTemplate renders a template from the cache or creates a new one if not present
func RenderTemplate(w http.ResponseWriter, t string) {
	tmpl, inMap := tc[t]
	if !inMap {
		// Template not in cache, create it
		log.Println("creating template and adding to cache")
		err := createTemplate(t)
		if err != nil {
			log.Println("error creating template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tmpl = tc[t]
	}

	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Println("error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// createTemplate parses the template files and adds them to the cache
func createTemplate(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// Parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// Add template to cache
	tc[t] = tmpl

	return nil
}
