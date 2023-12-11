package main

import (
    "database/sql"
	_ "github.com/lib/pq"
	"github.com/Tibirlayn/Go/handlers/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := sql.Open("postgres", "postgres://Server:admin@localhost/Account?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

    r := gin.Default()

	// Вызов обработчика маршрута
	routes.SetupRoutes(r)

    // Запуск сервера
    r.Run(":8080")
}
