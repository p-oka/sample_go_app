package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type templateHandler struct {
	once     sync.Once
	filename string
	db       *sql.DB
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates",
				t.filename)))
	})

	rows, err := t.db.Query("SELECT * FROM comments")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var comments []string

	for rows.Next() {
		var id int
		var comment string
		if err := rows.Scan(&id, &comment); err != nil {
			log.Fatal(err)
		} else {
			comments = append(comments, comment)
		}
	}

	t.templ.Execute(w, comments)
}

func main() {
	db, err := sql.Open("mysql", "mysql:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	http.Handle("/", &templateHandler{filename: "index.html", db: db})

	log.Println("Golang application starting on http://localhost:8080")
	log.Println("Ctrl-C to shutdown server")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
