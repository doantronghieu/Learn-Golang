package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
	now := time.Now().UTC()

	item := TodoItem{
		Id:          1,
		Title:       "Task 1",
		Description: "Content 1",
		Status:      "Doing",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	// --------------------------------------------------------------------------
	jsData, err := json.Marshal(item)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(jsData))
	// --------------------------------------------------------------------------
	jsString := "{\"id\":1,\"title\":\"Task 1\",\"description\":\"Content 1\",\"status\":\"Doing\",\"created_at\":\"2024-02-19T22:05:55.420951124Z\",\"updated_at\":\"2024-02-19T22:05:55.420951124Z\"}"
	var item2 TodoItem
	if err := json.Unmarshal([]byte(jsString), &item2); err != nil {
		log.Fatalln(err)
	}
	log.Println(item2)
	// --------------------------------------------------------------------------
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

	// Run the application on port 8081, listen on all available network interfaces (0.0.0.0).
	r.Run(":8081")
}
