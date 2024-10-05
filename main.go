package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id    int
	Title string
	Body  template.HTML
}

var db, err = sql.Open("mysql", "root:@tcp(localhost)/go_blog")

func main() {
	// stmt, err := db.Prepare("Insert into posts (title, body) values(?,?)")
	// checkErr(err)
	// _, err = stmt.Exec("My First post", "<p>This is my first content</p>")
	// checkErr(err)
	
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

func checkErr(err error){
	if err != nil {
		panic(err)
	}
}