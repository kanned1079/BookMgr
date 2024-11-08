package user

import (
	"book-mgr-backend/dao"
	"book-mgr-backend/handler"
	"book-mgr-backend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

func HandleGetSummary_User(context *gin.Context) {
	id, err := strconv.ParseInt(context.Query("user_id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	log.Println(id)
	if id <= 0 {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "提供的信息无效",
		})
		return
	}

	// 返回的数据
	responseData := &struct {
		Unreturned     int64   `json:"unreturned"`
		BorrowedNums   int64   `json:"borrowed_nums"`
		RankingPercent float64 `json:"ranking_percent"`
	}{}

	// 查询未归还书的数量
	if err := dao.Db.Model(&model.History{}).Where("user_id = ? AND is_back = ?", id, false).Count(&responseData.Unreturned).Error; err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	// 查询用户借阅的所有数量
	if err := dao.Db.Model(&model.History{}).Where("user_id = ?", id).Count(&responseData.BorrowedNums).Error; err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	// 查询 History 表中的所有数据量
	var totalBorrowed int64
	if err := dao.Db.Model(&model.History{}).Count(&totalBorrowed).Error; err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	// 计算借阅排名百分比
	if totalBorrowed > 0 {
		responseData.RankingPercent = float64(responseData.BorrowedNums) / float64(totalBorrowed) * 100
	} else {
		responseData.RankingPercent = 0
	}

	// 成功
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"summary": responseData,
		"msg":     "success",
	})
}

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
	Id        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	BookId    int64  `json:"book_id"`
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
	userId, _ := strconv.Atoi(c.DefaultQuery("user_id", "0")) // 获取 user_id

	// 如果 user_id 为 0，返回错误
	if userId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少用户ID"})
		return
	}

	// 计算偏移量
	offset := (page - 1) * size

	// 查询总记录数（不需要 Offset 和 Limit）
	var totalCount int64
	countQuery := db.Model(&model.History{}).Where("user_id = ?", userId)
	if name != "" {
		countQuery = countQuery.Joins("JOIN t_books ON t_books.id = t_history.book_id").
			Where("t_books.name LIKE ?", "%"+name+"%")
	}
	countQuery.Count(&totalCount)

	// 查询分页记录并预加载关联的书籍信息
	var histories []model.History
	query := db.Preload("Book").Where("user_id = ?", userId)
	if name != "" {
		query = query.Joins("JOIN t_books ON t_books.id = t_history.book_id").
			Where("t_books.name LIKE ?", "%"+name+"%")
	}
	if err := query.Offset(offset).Limit(size).Order("t_history.created_at DESC").Find(&histories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}

	// 构造响应
	var response []BorrowHistoryResponse
	for _, history := range histories {
		book := history.Book // 直接获取预加载的 Book 信息

		// 计算留存时间
		keepDuration := time.Since(*history.BorrowedAt)
		keep := keepDuration.String()

		// 构建响应数据
		response = append(response, BorrowHistoryResponse{
			Id:        history.Id,
			UserId:    history.UserId,
			BookId:    history.BookId,
			BorrowId:  history.BorrowId,
			CreatedAt: history.BorrowedAt.Format("2006-01-02"),
			IsBack:    history.IsBack,
			Keep:      keep,
			Name:      book.Name,
			ISBN:      book.ISBN,
		})
	}

	// 计算总页数
	pageCount := (totalCount + int64(size) - 1) / int64(size)

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"code":       200,
		"histories":  response,
		"page_count": pageCount,
	})
}

func HandleBorrowBookById_User(context *gin.Context) {
	postData := &struct {
		UserId int64 `json:"user_id"`
		BookId int64 `json:"book_id"`
	}{}
	if err := context.ShouldBind(postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "请提供用户和图书信息",
		})
		return
	}
	//log.Println(postData)

	if postData.UserId <= 0 || postData.BookId <= 0 {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "请求参数不正确",
		})
	}
	// 开启事务
	tx := dao.Db.Begin()

	// 检查图书库存
	var book model.Book
	if err := tx.First(&book, postData.BookId).Error; err != nil {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "无法查询图书",
		})
		return
	}
	if book.Residue <= 0 {
		tx.Rollback()
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "剩余数量不足",
		})
		return
	}

	// 将图书库存减1
	if err := tx.Model(&model.Book{}).Where("id = ?", postData.BookId).Update("residue", book.Residue-1).Error; err != nil {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "更新图书库存失败",
		})
		return
	}

	// 创建借阅订单号
	borrowId := createBorrowId(postData.UserId, postData.BookId)

	// 增加借阅记录
	nowTime := time.Now()
	history := model.History{
		BorrowId:   borrowId,
		UserId:     postData.UserId,
		BookId:     postData.BookId,
		BorrowedAt: &nowTime,
		IsBack:     false, // 未归还
	}
	if err := tx.Create(&history).Error; err != nil {
		tx.Rollback()
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "创建借阅记录失败",
		})
		return
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "提交事务失败",
		})
		return
	}

	// 成功响应
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "成功",
	})
}

func HandleReturnBookById_User(context *gin.Context) {
	postData := &struct {
		BorrowId string `json:"borrow_id"`
		UserId   int64  `json:"user_id"`
		BookId   int64  `json:"book_id"`
	}{}

	if err := context.ShouldBind(postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "请求参数错误",
		})
		return
	}

	if postData.UserId <= 0 || postData.BookId <= 0 || postData.BorrowId == "" {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "请求参数错误",
		})
		return
	}

	tx := dao.Db.Begin()

	// 更新 History 表，将 is_back 设置为 true
	if err := tx.Model(&model.History{}).
		Where("borrow_id = ? AND user_id = ? AND book_id = ?", postData.BorrowId, postData.UserId, postData.BookId).
		Update("is_back", true).Error; err != nil {
		tx.Rollback()
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "更新借阅记录失败",
		})
		return
	}

	// 更新 Book 表，将 residue +1
	if err := tx.Model(&model.Book{}).
		Where("id = ?", postData.BookId).
		Update("residue", gorm.Expr("residue + ?", 1)).Error; err != nil {
		tx.Rollback()
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "更新书籍库存失败",
		})
		return
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "事务提交失败",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "归还成功",
	})
}

// 创建订单号函数，格式为 <年><月><日><用户id><图书id><当天时间戳>
func createBorrowId(userId, bookId int64) string {
	timestamp := time.Now().Format("20060102") // 格式化日期
	return fmt.Sprintf("%s%d%d%d", timestamp, userId, bookId, time.Now().Unix()%86400)
}
