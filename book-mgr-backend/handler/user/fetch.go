package user

import (
	"book-mgr-backend/dao"
	"book-mgr-backend/handler"
	"book-mgr-backend/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func HandleGetAllBooks_User(context *gin.Context) {
	// 获取分页参数
	// 从查询参数中获取页码和页面大小，并检查是否有错误
	err, page, size := handler.GetPage2SizeFormQueryParams(context)
	if err != nil {
		// 如果缺少分页参数，则返回错误信息
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "缺少查询参数",
		})
		return
	}
	log.Println("Page:", page, "Size:", size)

	// 获取搜索和排序参数
	// search_by: 要搜索的字段（如 name, author, publisher），对应数据库中的列名
	searchBy := context.Query("search_by")
	// search_content: 搜索内容，即实际的搜索值
	searchContent := context.Query("search_content")
	// search_sort: 排序方式，如 "ASC" 或 "DESC"
	searchSort := context.Query("search_sort")

	// 获取总记录数
	// 查询书籍的总记录数，以便用于计算分页
	var totalBooks int64
	if result := dao.Db.Model(&model.Book{}).Count(&totalBooks); result.Error != nil {
		// 如果查询总记录数出错，则返回错误信息
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "查询总记录数出错",
		})
		return
	}

	// 计算总页数
	// 根据总记录数和页面大小计算总页数，用于返回给客户端
	pageCount := (totalBooks + int64(size) - 1) / int64(size)

	// 查询分页数据
	// 根据分页参数计算偏移量，并查询数据库中的书籍信息
	var books []model.Book
	offset := (page - 1) * size
	query := dao.Db.Model(&model.Book{}).Offset(int(offset)).Limit(int(size))

	// 如果指定了搜索条件，则添加过滤条件
	if searchBy != "" && searchContent != "" {
		query = query.Where(searchBy+" LIKE ?", "%"+searchContent+"%")
	}

	// 如果指定了排序方式，则添加排序条件
	if searchSort != "" {
		query = query.Order("`" + searchBy + "` " + searchSort)
	}

	// 执行查询
	if result := query.Find(&books); result.Error != nil {
		// 如果查询出错，则返回错误信息
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "查询出错",
		})
		return
	}

	// 返回分页数据和总页数
	// 返回书籍列表、总页数和总记录数给客户端
	context.JSON(http.StatusOK, gin.H{
		"code":        http.StatusOK,
		"books":       books,
		"page_count":  pageCount,
		"total_books": totalBooks,
	})
}

type BorrowHistoryResponse struct {
	Id        int    `json:"id"`
	BorrowId  string `json:"borrow_id"`
	CreatedAt string `json:"created_at"`
	IsBack    bool   `json:"is_back"`
	Keep      string `json:"keep"` // 留存时间
	Name      string `json:"name"`
	ISBN      string `json:"isbn"`
}

func HandleGetAllMyBorrowed_User(c *gin.Context) {
	// 从上下文中获取数据库实例
	db := dao.Db

	// 获取分页和查询参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	name := c.DefaultQuery("name", "")
	userId, _ := strconv.Atoi(c.DefaultQuery("user_id", "0")) // 获取user_id

	// 如果 user_id 为0，返回错误
	if userId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少用户ID"})
		return
	}

	// 计算偏移量
	offset := (page - 1) * size

	// 查询记录
	var histories []model.History
	query := db.Where("user_id = ?", userId)
	if name != "" {
		query = query.Joins("JOIN t_books ON t_books.id = t_history.book_id").
			Where("t_books.name LIKE ?", "%"+name+"%")
	}
	if err := query.Offset(offset).Limit(size).Order("created_at DESC").Find(&histories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}

	// 获取书籍信息并构造响应
	var response []BorrowHistoryResponse
	for _, history := range histories {
		var book model.Book
		if err := db.First(&book, history.BookId).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "书籍信息查询失败"})
			return
		}

		// 计算留存时间
		keepDuration := time.Since(*history.BorrowedAt)
		keep := keepDuration.String()

		// 构建响应数据
		response = append(response, BorrowHistoryResponse{
			Id:        history.Id,
			BorrowId:  history.BorrowId,
			CreatedAt: history.BorrowedAt.Format("2006-01-02"),
			IsBack:    history.IsBack,
			Keep:      keep,
			Name:      book.Name,
			ISBN:      book.ISBN,
		})
	}

	// 查询总记录数用于分页
	var totalCount int64
	query.Count(&totalCount)

	c.JSON(http.StatusOK, gin.H{
		"code":       200,
		"histories":  response,
		"page_count": (totalCount + int64(size) - 1) / int64(size),
	})
}
