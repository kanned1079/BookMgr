const { Sequelize, DataTypes } = require('sequelize');
const sequelize = require('../db/sequelize'); // 导入 Sequelize 实例

const User = sequelize.define(
    'User',
    {
        id: {
            type: DataTypes.BIGINT,
            primaryKey: true,
            autoIncrement: true,
        },
        role: {
            type: DataTypes.STRING,
            allowNull: false,
        },
        email: {
            type: DataTypes.STRING,
            allowNull: false,
            unique: true, // 确保邮箱唯一
            validate: {
                isEmail: true, // 验证邮箱格式
            },
        },
        password: {
            type: DataTypes.STRING,
            allowNull: false,
        },
    },
    {
        tableName: 'x_user', // 显式指定表名
        timestamps: true, // 启用 createdAt 和 updatedAt
        paranoid: true, // 启用 deletedAt（软删除）
        underscored: true
    }
);

module.exports = User;