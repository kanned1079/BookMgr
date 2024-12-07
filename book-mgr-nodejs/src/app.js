const express = require('express');
const {initializeDatabase, autoMigrateModels} = require('./db/init'); // 导入 Sequelize 实例
const sequelize = require('./db/sequelize');
const User = require('./model/user');

const {handleUserLogin, handleUserRegister} = require('./handler/auth');

const app = express();
const PORT = process.env.PORT || 3000;

// 初始化应用程序
async function initApp() {
    await initializeDatabase();  // 初始化数据库

    console.log('---------')

    // await autoMigrateModels();

    app.use(express.json()); // 中间件：解析 JSON 请求体

    // 示例路由
    app.get('/', (req, res) => {
        res.send('Server is running and database is connected!');
    });

    /*
    * adminGroup := r.Group("/api/admin/v1")
	{
		adminGroup.POST("login", univer.HandleUserLogin)

		adminGroup.GET("summary", admin.GetAdminSummary_Admin)

		adminGroup.GET("book", admin.HandleGetAllBooks_Admin)
		adminGroup.POST("book", admin.HandleAddBook_Admin)
		adminGroup.PUT("book", admin.HandleUpdateBook_Admin)
		adminGroup.DELETE("book", admin.HandleDeleteBook_Admin)

		adminGroup.GET("user", admin.HandleGetAllUsers_Admin)

		adminGroup.GET("history", admin.GetAllHistories_Admin)
	}

	userGroup := r.Group("/api/user/v1")
	{
		userGroup.POST("login", univer.HandleUserLogin)
		userGroup.POST("register", univer.HandleUserRegister)
		userGroup.GET("summary", user.HandleGetSummary_User)
		userGroup.GET("book", user.HandleGetAllBooks_User)
		userGroup.GET("history", user.HandleGetAllMyBorrowed_User)
		userGroup.PATCH("history", user.HandleReturnBookById_User)
		userGroup.POST("borrow", user.HandleBorrowBookById_User)
	}
    * */

    // 管理员
    app.post('api/admin/v1/login', handleUserLogin)


    // 用户
    app.post('api/user/v1/register', handleUserRegister)



    // 启动服务器
    app.listen(PORT, () => {
        console.log(`Server is running on http://localhost:${PORT}`);
    });
}

// 启动应用程序
initApp();