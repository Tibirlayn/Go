package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	    // Обработчик маршрута
		r.GET("/hello", func(c *gin.Context) {
			c.String(200, "Hello, World!")
		})

		// Обработчик маршрута
		r.GET("/game", func(c *gin.Context) {
			c.String(200, "Hello, Game!")
		})

		// Обработчик маршрута
		r.GET("/account", func(c *gin.Context) {
			c.String(200, "Hello, Game!")
		})

		// Обработчик маршрута
		r.GET("/mail", func(c *gin.Context) {
			c.String(200, "Hello, Game!")
		})

}