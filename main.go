package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Подключение к базе данных (здесь используется PostgreSQL, но можно выбрать другую базу данных)
	db, err := sql.Open("postgres", "user=username dbname=youdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	
}
