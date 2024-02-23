package main

import (
	// "encoding/json"
	// "log"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"social-todo-list/common"
	"social-todo-list/modules/item/model"
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

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", ginitem.CreateItem(db))
			items.GET("", ListItem(db))
			items.GET("/:id", ginitem.GetItem(db))
			items.PATCH("/:id", UpdateItem(db))
			items.DELETE("/:id", DeleteItem(db))
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

// UpdateItem is a handler function for updating a TodoItem using Gin framework
func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		// Parse the URL parameter to get the ID of the TodoItem to be updated
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			// Return a Bad Request response if there is an error in parsing the ID
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Create a TodoItemUpdate instance to hold the update data
		var updateData model.TodoItemUpdate

		// Bind the incoming JSON data to the TodoItemUpdate struct
		if err := c.ShouldBind(&updateData); err != nil {
			// Return a Bad Request response if there is an error in data binding
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Update the TodoItem in the database with the provided ID
		if err := db.Where("id = ?", id).Updates(&updateData).Error; err != nil {
			// Return a Bad Request response if there is an error in the database update
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

// DeleteItem is a handler function for soft deleting a TodoItem using Gin framework
func DeleteItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		// Parse the URL parameter to get the ID of the TodoItem to be deleted
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			// Return a Bad Request response if there is an error in parsing the ID
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Soft delete the TodoItem by updating its status to "Deleted" in the database
		if err := db.Table(model.TodoItem{}.TableName()).Where("id = ?", id).
			Updates(map[string]interface{}{"status": "Deleted"}).
			Error; err != nil {
			// Return a Bad Request response if there is an error in the database update
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

// ListItem is a handler function for listing TodoItems using Gin framework with pagination
func ListItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		// Create a common.Paging struct to hold pagination parameters
		var paging common.Paging

		// Bind the incoming JSON data to the Paging struct
		if err := c.ShouldBind(&paging); err != nil {
			// Return a Bad Request response if there is an error in data binding
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Process pagination parameters
		paging.Process()

		// Initialize a slice to hold the resulting TodoItems
		var result []model.TodoItem

		if err := db.Table(model.TodoItem{}.TableName()).
			// Exclude soft-deleted TodoItems by filtering on status
			Where("status <> ?", "Deleted").
			Select("id").
			// Retrieve the total count of TodoItems
			Count(&paging.Total).
			Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Retrieve a paginated list of TodoItems ordered by ID in descending order
		if err := db.
			Select("*").
			Order("id desc").
			Offset((paging.Page - 1) * paging.Limit).
			Limit(paging.Limit).
			Find(&result).
			Error; err != nil {
			// Return a Bad Request response if there is an error in retrieving the list
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
