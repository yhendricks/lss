package main

import (
	"net/http"
	"html/template"
	"log"
	"os"
	"io/ioutil"
	//"fmt"
	"github.com/yhendricks/lss/webapp/viewmodel"
)

func main() {
	templates := populateTemplate()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]			// strip off first char
		t := templates[requestedFile + ".html"]
		var context interface{}
		switch requestedFile {
		case "shop":
			context =  viewmodel.NewShop()
		default:
			context = viewmodel.NewBase()
		}
		if t != nil {
			err := t.Execute(w, context)
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

func populateTemplate() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "src/github.com/yhendricks/lss/templates"
	//const basePath = "src/github.com/yhendricks/lss/templates"
	//fmt.Println(basePath + "/_layout.html")
	layout := template.Must(template.ParseFiles(basePath +  "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))
	dir, err := os.Open(basePath+"/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of the content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath+"/content/"+fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read contents from file '" + fi.Name()+ "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name()+ "'")
		}
		result[fi.Name()] = tmpl
	}
	return result
}
