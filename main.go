package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	if err := t.templ.Execute(w, nil); err != nil {
		panic("Something bad happened")
	}
}

func main() {
	http.Handle("/", &templateHandler{
		filename: "chat.html",
	})

	addr := ":5050"

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}

}

