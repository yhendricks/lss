package main

import (
	"net/http"
	"html/template"
	"log"
)

func main() {
	templates := populateTemplate()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]			// strip off first char
		t := templates.Lookup(requestedFile + ".html")
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	http.Handle("/img/", http.FileServer(http.Dir("src/github.com/yhendricks/lss/public")))
	http.Handle("/css/", http.FileServer(http.Dir("src/github.com/yhendricks/lss/public")))
	http.ListenAndServe(":8000", nil)
}

func populateTemplate() *template.Template {
	result := template.New("templates")
	const basePath = "src/github.com/yhendricks/lss/templates"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}
