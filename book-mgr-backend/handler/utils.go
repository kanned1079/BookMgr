package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPage2SizeFormQueryParams(context *gin.Context) (err error, page int64, size int64) {
	page, err = strconv.ParseInt(context.DefaultQuery("page", "1"), 10, 64)
	size, err = strconv.ParseInt(context.DefaultQuery("size", "10"), 10, 64)
	return
}
