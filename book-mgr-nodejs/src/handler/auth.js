const { User } = require('../model/user'); // 导入 Sequelize 模型

async function handleUserLogin(req, res) {
    const { email, password, role } = req.body;

    // 检查请求数据是否合法
    if (!email || !password || !role) {
        return res.status(400).json({
            code: 400,
            authed: false,
            msg: '请求参数错误',
        });
    }

    try {
        // 查找用户
        const user = await User.findOne({ where: { email } });

        if (!user) {
            return res.status(404).json({
                code: 404,
                authed: false,
                msg: '用户不存在，请注册',
            });
        }

        // 检查角色
        if (user.role !== role) {
            return res.status(403).json({
                code: 403,
                authed: false,
                msg: '非法访问',
            });
        }

        // 检查密码
        if (user.password !== password) {
            return res.status(401).json({
                code: 401,
                authed: false,
                msg: '密码错误',
            });
        }

        // 隐藏密码
        user.password = undefined;

        return res.status(200).json({
            code: 200,
            authed: true,
            user,
            msg: '验证通过',
        });
    } catch (error) {
        console.error('Error during user login:', error);
        return res.status(500).json({
            code: 500,
            authed: false,
            msg: '服务器错误，请稍后重试',
        });
    }
}

async function handleUserRegister(req, res) {
    const { email, password } = req.body;

    // 检查请求数据是否合法
    if (!email || !password) {
        return res.status(400).json({
            code: 400,
            msg: '请求数据无效',
        });
    }

    try {
        // 启动事务
        const result = await sequelize.transaction(async (transaction) => {
            // 检查用户是否已存在
            const existingUser = await User.findOne({ where: { email }, transaction });

            if (existingUser) {
                throw { code: 409, msg: '用户已存在' };
            }

            // 创建新用户
            return await User.create(
                {
                    email,
                    password, // 在实际应用中应对密码进行哈希处理
                    role: 'user',
                },
                {transaction}
            );
        });

        return res.status(200).json({
            code: 200,
            registered: true,
            msg: '注册成功',
            user: result,
        });
    } catch (error) {
        if (error.code === 409) {
            return res.status(409).json({
                code: 409,
                registered: false,
                msg: error.msg,
            });
        }

        console.error('Error during user registration:', error);
        return res.status(500).json({
            code: 500,
            registered: false,
            msg: '服务器错误，请稍后重试',
        });
    }
}

module.exports = {handleUserLogin, handleUserRegister}