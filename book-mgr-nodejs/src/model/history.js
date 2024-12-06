const { Sequelize, DataTypes } = require('sequelize');
const sequelize = require('../db/sequelize'); // 导入 Sequelize 实例
const Book = require('./book'); // 导入 Book 模型

const History = sequelize.define(
    'History',
    {
        id: {
            type: DataTypes.BIGINT,
            primaryKey: true,
            autoIncrement: true,
        },
        borrow_id: {
            type: DataTypes.STRING,
            allowNull: false,
        },
        user_id: {
            type: DataTypes.BIGINT,
            allowNull: false,
        },
        book_id: {
            type: DataTypes.BIGINT,
            allowNull: false,
        },
        borrowed_at: {
            type: DataTypes.DATE,
            allowNull: true,
        },
        is_back: {
            type: DataTypes.BOOLEAN,
            allowNull: false,
            defaultValue: false, // 默认值为 false
        },
    },
    {
        tableName: 'x_history', // 显式指定表名
        timestamps: true, // 启用 createdAt 和 updatedAt
        paranoid: true, // 启用 deletedAt（软删除）
        underscored: true
    }
);

// 外键关联
History.belongsTo(Book, {
    foreignKey: 'book_id', // 指定外键
    as: 'book', // 设置关联的别名
});

module.exports = History;