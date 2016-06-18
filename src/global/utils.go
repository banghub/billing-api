package global

import (
	"fmt"
	"html/template"
	"net/http"
)

// Response : data type responses
type Response struct {
	Links  interface{} `json:"links"`
	Data   interface{} `json:"data"`
	Errors interface{} `json:"errors"`
}

// Index : index page
func Index(w http.ResponseWriter, r *http.Request) {
	var test interface{}
	RenderTemplate(w, "index", test)
}

// NotFound : handle 404.
func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "custom 404")
}

// RenderTemplate : render template
func RenderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	filePath := tmpl + ".html"
	t, err := template.ParseFiles(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
