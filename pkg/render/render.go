package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// renders template using html template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}

var tc = make(map[string]*template.Template)
function RenderTemplateTest(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//check if we already have template in cache
	_, inMap := tc[t]
	if !inMap {
		// need to create the template 
		err = createTemplateCache(t)
		if err != 
	} else {
		log.Println("Using cached template")
	}

	tmpl := tc[t]

	err = tmpl.Execute(w, nill)
}

func createTemplate(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil () {
		return err
	}

	// add template to cache
	tc[t] = tmpl

	return nil
}