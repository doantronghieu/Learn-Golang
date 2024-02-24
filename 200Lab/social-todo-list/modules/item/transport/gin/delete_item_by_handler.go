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

func DeleteItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewDeleteItemBiz(store)

		if err := business.DeleteItemById(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
