package ginitem

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"social-todo-list/common"
	"social-todo-list/modules/item/biz"
	"social-todo-list/modules/item/storage"
)

// GetItem returns a Gin handler function that fetches an item from the database
//  using the provided *gorm.DB.
func GetItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		// Extract item ID from the request parameters
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			// Handle invalid ID error
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Create a new SQL store and business logic instance
		store := storage.NewSQLStore(db)
		business := biz.NewGetItemBiz(store)

		// Fetch the item using the business logic
		data, err := business.GetItemById(c.Request.Context(), id)
		if err != nil {
			// Handle error during item retrieval
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Return successful response with the retrieved item data
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
