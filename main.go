package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		title := vars["title"]
// 		page := vars["page"]
// 		fmt.Fprintf(w, "You`ve requested the book:%s on Page: %s\n", title, page)
// 	}).Methods("GET")
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello ,you`ve requested:%s", r.URL.Path)
// 	})
// 	http.ListenAndServe(":80", r)
// }

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() { log.Println(r.URL.Path, time.Since(start)) }()
		f(w, r)

	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	data := TodoPageData{
		PageTitle: "My Todo List",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
		},
	}
	tmpl.Execute(w, data)
}

func main() {

	http.HandleFunc("/", logging(HomeHandler))
	http.ListenAndServe(":80", nil)
}
