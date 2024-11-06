package admin

import (
	"book-mgr-backend/dao"
	"book-mgr-backend/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandleAdminLogin(context *gin.Context) {
	postData := &struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := context.ShouldBind(postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
	}
	log.Println(postData)

	var user model.User
	if result := dao.Db.Model(&model.User{}).Where("email = ?", postData.Email).First(&user); result.RowsAffected == 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":   http.StatusNotFound,
			"authed": false,
			"msg":    "用户不存在",
		})
		return
	}
	if user.Password != postData.Password {
		context.JSON(http.StatusOK, gin.H{
			"code":   http.StatusUnauthorized,
			"authed": false,
			"msg":    "密码不正确",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"authed": true,
		"msg":    "登录成功",
	})
}
