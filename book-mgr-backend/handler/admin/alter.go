package admin

import (
	"book-mgr-backend/dao"
	"book-mgr-backend/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func HandleAddBook_Admin(context *gin.Context) {
	postData := &struct {
		Name      string  `json:"name"`
		Publisher string  `json:"publisher"`
		Year      int32   `json:"year"`
		Remark    string  `json:"remark"`
		Author    string  `json:"author"`
		ISBN      string  `json:"isbn"`
		Price     float64 `json:"price"`
		Residue   int64   `json:"residue"`
	}{}
	if err := context.ShouldBind(postData); err != nil {
		log.Println(err.Error())
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"created": false,
		})
		return
	}
	log.Println(postData)
	var newBook = model.Book{
		Name:      postData.Name,
		Publisher: postData.Publisher,
		Year:      postData.Year,
		Remark:    postData.Remark,
		Author:    postData.Author,
		ISBN:      postData.ISBN,
		Price:     postData.Price,
		Residue:   postData.Residue,
	}
	if err := dao.Db.Model(&model.Book{}).Create(&newBook).Error; err != nil {
		log.Println(err.Error())
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"created": false,
			"msg":     err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"created": true,
	})
}

func HandleUpdateBook_Admin(context *gin.Context) {
	log.Println("修改图书信息")
	postData := &struct {
		Id        int     `json:"id"`
		Name      string  `json:"name"`
		Publisher string  `json:"publisher"`
		Year      int32   `json:"year"`
		Remark    string  `json:"remark"`
		Author    string  `json:"author"`
		ISBN      string  `json:"isbn"`
		Price     float64 `json:"price"`
		Residue   int64   `json:"residue"`
	}{}
	if err := context.ShouldBind(&postData); err != nil {
		log.Println(err.Error())
	}
	log.Println(postData)
	// 查找并更新数据库中的记录
	var book model.Book
	if err := dao.Db.Model(&model.Book{}).Where("id = ?", postData.Id).First(&book).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"msg":     "书籍未找到",
			"updated": false,
		})
		return
	}

	// 更新字段
	book.Name = postData.Name
	book.Publisher = postData.Publisher
	book.Year = postData.Year
	book.Remark = postData.Remark
	book.Author = postData.Author
	book.ISBN = postData.ISBN
	book.Price = postData.Price
	book.Residue = postData.Residue

	// 保存更改
	if err := dao.Db.Save(&book).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"msg":     "更新书籍信息失败",
			"updated": false,
		})
		return
	}

	// 返回更新后的书籍信息
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"msg":     "书籍信息更新成功",
		"updated": true,
	})
}

func HandleDeleteBook_Admin(context *gin.Context) {
	bookId, err := strconv.ParseInt(context.Query("id"), 10, 64)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(bookId)
	if result := dao.Db.Model(&model.Book{}).Where("id = ?", bookId).Delete(&model.Book{}); result.RowsAffected == 0 {
		log.Println("删除失败")
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"deleted": true,
	})
}
