package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Post struct {
	Id    int
	Title string
	Body  template.HTML
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		post := Post{Id: 1, Title: "First Post", Body: template.HTML("<strong>This is my first Post!</strong>")}

		t := template.Must(template.ParseFiles("templates/index.html"))
		executeTemplate := t.ExecuteTemplate(w, "index.html", post)

		if err := executeTemplate; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
