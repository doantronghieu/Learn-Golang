package ginitem

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"social-todo-list/common"
	"social-todo-list/modules/item/biz"
	"social-todo-list/modules/item/model"
	"social-todo-list/modules/item/storage"
)

func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		var updateData model.TodoItemUpdate

		if err := c.ShouldBind(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewUpdateItemBiz(store)

		if err := business.UpdateItemById(c.Request.Context(), id, &updateData); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
