const { Sequelize, DataTypes } = require('sequelize');
const sequelize = require('../db/sequelize');  // 导入 Sequelize 实例

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
    paranoid: true,
    underscored: true
});

module.exports = Book;