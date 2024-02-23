package ginitem

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"social-todo-list/common"
	"social-todo-list/modules/item/biz"
	"social-todo-list/modules/item/model"
	"social-todo-list/modules/item/storage"
)

// Handler function for creating a new TodoItem using Gin framework
func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		// Parse the incoming JSON data into TodoItemCreation struct
		var data model.TodoItemCreation
		
		if err := c.ShouldBind(&data); err != nil {
			// Return a Bad Request response if there is an error in data binding
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Create a new SQL data store and business logic instance
		store := storage.NewSQLStore(db)
		business := biz.NewCreateItemBiz(store)

		// Call the business logic to create a new TodoItem
		if err := business.CreateNewItem(c.Request.Context(), &data); err != nil {
			// Return a Bad Request response if there is an error in business logic
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Return a success response with the new TodoItem's ID
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
