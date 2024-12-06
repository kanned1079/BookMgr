const sequelize = require('./sequelize')
// 导入模型
const User = require('../model/user');
const Book = require('../model/book');
const History = require('../model/history');

async function initializeDatabase() {
    console.log("开始初始化数据库连接", process.env.DB_NAME, process.env.DB_HOST);
    try {
        await sequelize.authenticate().then(result => {
        })
        console.log("开始自动迁移")
        await autoMigrateModels()
        console.log('Database connection has been established successfully.');
        // await sequelize.sync({ alter: true });
        // console.log('All models were synchronized successfully.');

        // await autoMigrateModels()
    } catch (error) {
        console.error('Database initialization failed:', error.message);
        process.exit(1); // 如果初始化失败，退出程序
    }
}

async function autoMigrateModels() {
    try {
        await sequelize.sync({ alter: true }) // 自动迁移模型
        console.log('All models were synchronized successfully.');
    } catch (error) {
        console.error('Failed to synchronize models:', error.message);
        process.exit(1); // 如果迁移失败，退出程序
    }
}



module.exports = {initializeDatabase, autoMigrateModels};