const express = require('express');
const {initializeDatabase, autoMigrateModels} = require('./db/init'); // 导入 Sequelize 实例
const sequelize = require('./db/sequelize');
const User = require('./model/user');

const app = express();
const PORT = process.env.PORT || 3000;

// 初始化应用程序
async function initApp() {
    await initializeDatabase();  // 初始化数据库

    console.log('---------')
    let testUser = new User({
        email: 'test@test.com',
    })

    const user1 = await User.create({
        email: 'test@test.com',
        password: 'test123',
        role: 'admin',
    })

    // await autoMigrateModels();

    app.use(express.json()); // 中间件：解析 JSON 请求体

    // 示例路由
    app.get('/', (req, res) => {
        res.send('Server is running and database is connected!');
    });

    // 启动服务器
    app.listen(PORT, () => {
        console.log(`Server is running on http://localhost:${PORT}`);
    });
}

// 启动应用程序
initApp();