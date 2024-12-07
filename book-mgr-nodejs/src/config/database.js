const { Sequelize, DataTypes} = require('sequelize');
// const User = require('../model/user');
// const Book = require('../model/book'); // 导入模型
// const History = require('../model/history');
require('dotenv').config({
    path: './config/.env',
});

const sequelize = new Sequelize(
    process.env.DB_NAME,
    process.env.DB_USER,
    process.env.DB_PASSWORD,
    {
        host: process.env.DB_HOST,
        port: process.env.DB_PORT,
        dialect: 'mysql',
        logging: false,
    }
);

const Book = sequelize.define(
    'Book',
    {
        id: {
            type: DataTypes.BIGINT,
            primaryKey: true,
            autoIncrement: true,
        },
        name: {
            type: DataTypes.STRING,
            allowNull: false,
        },
        publisher: {
            type: DataTypes.STRING,
            allowNull: true,
        },
        year: {
            type: DataTypes.INTEGER,
            allowNull: true,
        },
        remark: {
            type: DataTypes.TEXT,
            allowNull: true,
        },
        author: {
            type: DataTypes.STRING,
            allowNull: true,
        },
        isbn: {
            type: DataTypes.STRING,
            allowNull: true,
        },
        price: {
            type: DataTypes.FLOAT,
            allowNull: true,
        },
        residue: {
            type: DataTypes.BIGINT,
            allowNull: true,
        },
        cover_url: {
            type: DataTypes.TEXT,
            allowNull: true,
        },
    }, {
        tableName: 'x_books', // 显式指定表名
        timestamps: true, // 启用 createdAt 和 updatedAt
        paranoid: true, // 启用 deletedAt
    });



async function autoMigrateModels() {
    try {
        await sequelize.sync({ alter: true }); // 自动迁移模型
        console.log('All models were synchronized successfully.');
    } catch (error) {
        console.error('Failed to synchronize models:', error.message);
        process.exit(1); // 如果迁移失败，退出程序
    }
}

module.exports = { sequelize, initializeDatabase, autoMigrateModels, Book };