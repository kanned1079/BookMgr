package univer

import (
	"book-mgr-backend/dao"
	"book-mgr-backend/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func HandleUserLogin(context *gin.Context) {
	postData := &struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}{}
	if err := context.ShouldBindJSON(postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":   http.StatusBadRequest,
			"authed": false,
			"msg":    "请求参数错误",
		})
		return
	}

	var user model.User
	if err := dao.Db.Where("email = ?", postData.Email).First(&user).Error; err != nil {
		// 用户不存在
		context.JSON(http.StatusOK, gin.H{
			"code":   http.StatusNotFound,
			"authed": false,
			"msg":    "用户不存在，请注册",
		})
		return
	}

	if user.Role != postData.Role {
		context.JSON(http.StatusOK, gin.H{
			"code":   http.StatusForbidden,
			"authed": false,
			"msg":    "非法访问",
		})
		return
	}

	if user.Password != postData.Password {
		context.JSON(http.StatusOK, gin.H{
			"code":   http.StatusUnauthorized,
			"authed": false,
			"msg":    "密码错误",
		})
		return
	}

	user.Password = ""
	context.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"authed": true,
		"user":   user,
		"msg":    "验证通过",
	})
}

func HandleUserRegister(context *gin.Context) {
	// 解析请求中的数据
	postData := &struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := context.ShouldBind(postData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "请求数据无效",
		})
		return
	}
	var newUser = model.User{
		Email:    postData.Email,
		Password: postData.Password, // 在真实应用中，建议先对密码进行哈希处理
		Role:     "user",
	}
	tx := dao.Db.Begin()
	var existingUser model.User
	if err := tx.Where("email = ?", newUser.Email).First(&existingUser).Error; err == nil {
		// 若找到匹配用户，说明用户已存在
		tx.Rollback() // 回滚事务
		context.JSON(http.StatusConflict, gin.H{
			"code":       http.StatusConflict,
			"registered": false,
			"msg":        "用户已存在",
		})
		return
	} else if err != gorm.ErrRecordNotFound {
		// 查询过程中发生错误
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "服务器错误，请稍后重试",
		})
		return
	}

	if err := tx.Create(&newUser).Error; err != nil {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":       http.StatusInternalServerError,
			"registered": false,
			"msg":        "注册失败",
		})
		return
	}

	tx.Commit()

	context.JSON(http.StatusOK, gin.H{
		"code":       http.StatusOK,
		"registered": true,
		"msg":        "注册成功",
	})
}
