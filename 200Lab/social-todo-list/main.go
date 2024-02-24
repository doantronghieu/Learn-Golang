package main

import (
	// "encoding/json"
	// "log"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"social-todo-list/common"
	"social-todo-list/middleware"
	ginitem "social-todo-list/modules/item/transport/gin"
)

func main() {
	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file")
	}

	// Get the database connection string from environment variable
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	// Check if the environment variable is set
	if dbConnectionString == "" {
		log.Fatalln("DB_CONNECTION_STRING environment variable is not set")
	}

	db, err := gorm.Open(mysql.Open(dbConnectionString), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	log.Println("DB Connection: ", db)

	// Create a new Gin router with default middleware (logger and recovery).
	r := gin.Default()
	r.Use(middleware.Recover())

	v1 := r.Group(
		"/v1",
		// middleware.Recovery(),
	)
	{
		items := v1.Group("/items")
		{
			items.POST("", ginitem.CreateItem(db))
			items.GET("", ginitem.ListItem(db))
			items.GET("/:id", ginitem.GetItem(db))
			items.PATCH("/:id", ginitem.UpdateItem(db))
			items.DELETE("/:id", ginitem.DeleteItem(db))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		go func() {
			defer common.Recovery()
			fmt.Println([]int{}[0])
		}()

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	port := os.Getenv("GIN_PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("Starting the application on port %s...", port)
	r.Run(":" + port)
}
