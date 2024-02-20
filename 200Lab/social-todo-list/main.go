package main

import (
	// "encoding/json"
	// "log"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

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

	fmt.Println(db)

	// now := time.Now().UTC()

	// item := TodoItem{
	// 	Id:          1,
	// 	Title:       "Task 1",
	// 	Description: "Content 1",
	// 	Status:      "Doing",
	// 	CreatedAt:   &now,
	// 	UpdatedAt:   &now,
	// }

	// Create a new Gin router with default middleware (logger and recovery).
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("")
			items.GET("")
			items.GET("/:id")
			items.PATCH("/:id")
			items.DELETE("/:id")
		}
	}

	// Define a route for handling HTTP GET requests to the "/ping" endpoint.
	r.GET("/ping", func(c *gin.Context) {
		// Respond with a JSON message indicating "pong" and an HTTP status OK (200).
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
