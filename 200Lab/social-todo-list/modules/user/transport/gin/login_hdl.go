package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"social-todo-list/common"
	"social-todo-list/component/tokenprovider"
	"social-todo-list/modules/user/biz"
	"social-todo-list/modules/user/model"
	"social-todo-list/modules/user/storage"
)

func Login(db *gorm.DB, tokenProvider tokenprovider.Provider) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData model.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		md5 := common.NewMd5Hash()

		business := biz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)

		account, err := business.Login(c.Request.Context(), &loginUserData)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
