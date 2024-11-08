package admin

import (
	"book-mgr-backend/dao"
	"book-mgr-backend/handler"
	"book-mgr-backend/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetAdminSummary_Admin(context *gin.Context) {
	responseData := &struct {
		UserCount     int64 `json:"user_count"`
		BookCount     int64 `json:"book_count"`
		BorrowedCount int64 `json:"borrowed_count"`
	}{}

	// 查询总藏书量
	if err := dao.Db.Model(&model.Book{}).Count(&responseData.BookCount).Error; err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	// 查询总用户数量
	if err := dao.Db.Model(&model.User{}).Count(&responseData.UserCount).Error; err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	// 查询总借阅量
	if err := dao.Db.Model(&model.History{}).Count(&responseData.BorrowedCount).Error; err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	// 成功
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"summary": responseData,
		"msg":     "success",
	})
}

func HandleGetAllBooks_Admin(context *gin.Context) {
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

func HandleGetAllUsers_Admin(context *gin.Context) {
	// 从请求参数中获取分页和筛选条件
	page := context.DefaultQuery("page", "1")
	size := context.DefaultQuery("size", "100")
	searchEmail := context.DefaultQuery("search_email", "")

	// 转换分页参数
	pageInt, _ := strconv.Atoi(page)
	sizeInt, _ := strconv.Atoi(size)

	// 创建查询链以筛选用户
	var users []model.User
	query := dao.Db.Model(&model.User{}).Where("deleted_at IS NULL")

	// 如果提供了邮箱搜索关键字，则添加筛选条件
	if searchEmail != "" {
		query = query.Where("email LIKE ?", "%"+searchEmail+"%")
	}

	// 分页查询用户
	err := query.Offset((pageInt - 1) * sizeInt).Limit(sizeInt).Find(&users).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "查询用户失败", "error": err.Error()})
		return
	}

	// 初始化用户列表
	type ResponseUser struct {
		Id           int64  `json:"id"`
		Role         string `json:"role"`
		Email        string `json:"email"`
		BorrowedNums int    `json:"borrowed_nums"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
	}

	var responseUsers []ResponseUser

	// 遍历用户查询未归还的书籍数量
	for _, user := range users {
		var borrowedNums int64
		// 查询 History 表中该用户未归还的记录数
		err := dao.Db.Model(&model.History{}).Where("user_id = ? AND is_back = ?", user.Id, false).Count(&borrowedNums).Error
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "查询借阅数量失败", "error": err.Error()})
			return
		}

		// 创建响应结构的用户数据
		responseUsers = append(responseUsers, ResponseUser{
			Id:           user.Id,
			Role:         user.Role,
			Email:        user.Email,
			BorrowedNums: int(borrowedNums),
			CreatedAt:    user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	// 计算总用户数，用于前端分页
	var totalUsers int64
	err = dao.Db.Model(&model.User{}).Where("deleted_at IS NULL").Count(&totalUsers).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "获取用户总数失败", "error": err.Error()})
		return
	}

	// 返回分页数据和用户列表
	context.JSON(http.StatusOK, gin.H{
		"code":       200,
		"users":      responseUsers,
		"page_count": (totalUsers + int64(sizeInt) - 1) / int64(sizeInt), // 计算页数
	})
}
