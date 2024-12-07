const { Book } = require('../model/book'); // 引入 Sequelize 模型
const { User } = require('../model/user');
const { History } = require('../model/history');

const handleAddBookAdmin = async (req, res) => {
    try {
        const postData = req.body;

        // 验证请求数据
        if (!postData.name || !postData.publisher) {
            return res.status(400).json({
                code: 400,
                created: false,
                msg: "缺少必要参数",
            });
        }

        // 创建新书记录
        const newBook = await Book.create({
            name: postData.name,
            publisher: postData.publisher,
            year: postData.year,
            remark: postData.remark,
            author: postData.author,
            isbn: postData.isbn,
            price: postData.price,
            residue: postData.residue,
            cover_url: postData.cover_url,
        });

        res.status(200).json({
            code: 200,
            created: true,
            book: newBook,
        });
    } catch (error) {
        console.error(error);
        res.status(500).json({
            code: 500,
            created: false,
            msg: error.message,
        });
    }
};

const handleUpdateBookAdmin = async (req, res) => {
    try {
        const { book_id, name, publisher, year, remark, author, isbn, price, residue, cover_url } = req.body;

        // 验证书籍 ID
        if (!book_id) {
            return res.status(400).json({
                code: 400,
                updated: false,
                msg: "缺少 book_id 参数",
            });
        }

        // 查找书籍
        const book = await Book.findByPk(book_id);
        if (!book) {
            return res.status(404).json({
                code: 404,
                updated: false,
                msg: "书籍未找到",
            });
        }

        // 更新字段
        book.name = name || book.name;
        book.publisher = publisher || book.publisher;
        book.year = year || book.year;
        book.remark = remark || book.remark;
        book.author = author || book.author;
        book.isbn = isbn || book.isbn;
        book.price = price || book.price;
        book.residue = residue || book.residue;
        book.cover_url = cover_url || book.cover_url;

        // 保存更新
        await book.save();

        res.status(200).json({
            code: 200,
            updated: true,
            book,
        });
    } catch (error) {
        console.error(error);
        res.status(500).json({
            code: 500,
            updated: false,
            msg: error.message,
        });
    }
};

const handleDeleteBookAdmin = async (req, res) => {
    try {
        const bookId = req.query.id;

        // 验证书籍 ID
        if (!bookId) {
            return res.status(400).json({
                code: 400,
                deleted: false,
                msg: "缺少 id 参数",
            });
        }

        // 删除书籍
        const result = await Book.destroy({
            where: { id: bookId },
        });

        if (result === 0) {
            return res.status(404).json({
                code: 404,
                deleted: false,
                msg: "书籍未找到",
            });
        }

        res.status(200).json({
            code: 200,
            deleted: true,
        });
    } catch (error) {
        console.error(error);
        res.status(500).json({
            code: 500,
            deleted: false,
            msg: error.message,
        });
    }
};

const getAdminSummaryAdmin = async (req, res) => {
    try {
        const responseData = {
            user_count: 0,
            book_count: 0,
            borrowed_count: 0,
        };

        responseData.book_count = await Book.count();
        responseData.user_count = await User.count();
        responseData.borrowed_count = await History.count();

        res.status(200).json({
            code: 200,
            summary: responseData,
            msg: 'success',
        });
    } catch (error) {
        console.error(error);
        res.status(500).json({
            code: 500,
            msg: error.message,
        });
    }
};

const handleGetAllBooksAdmin = async (req, res) => {
    try {
        const { page = 1, size = 10, search_by, search_content, search_sort = 'ASC' } = req.query;

        const pageNum = parseInt(page, 10);
        const pageSize = parseInt(size, 10);

        const where = search_by && search_content ? {
            [search_by]: { [Op.like]: `%${search_content}%` },
        } : {};

        const totalBooks = await Book.count({ where });
        const pageCount = Math.ceil(totalBooks / pageSize);

        const books = await Book.findAll({
            where,
            offset: (pageNum - 1) * pageSize,
            limit: pageSize,
            order: search_by ? [[search_by, search_sort]] : undefined,
        });

        res.status(200).json({
            code: 200,
            books,
            page_count: pageCount,
            total_books: totalBooks,
        });
    } catch (error) {
        console.error(error);
        res.status(500).json({
            code: 500,
            msg: '查询书籍失败',
            error: error.message,
        });
    }
};

const handleGetAllUsersAdmin = async (req, res) => {
    try {
        const { page = 1, size = 100, search_email = '' } = req.query;

        const pageNum = parseInt(page, 10);
        const pageSize = parseInt(size, 10);

        const where = search_email ? {
            email: { [Op.like]: `%${search_email}%` },
        } : {};

        const users = await User.findAll({
            where,
            offset: (pageNum - 1) * pageSize,
            limit: pageSize,
        });

        const responseUsers = [];

        for (const user of users) {
            const borrowedNums = await History.count({
                where: { user_id: user.id, is_back: false },
            });

            responseUsers.push({
                id: user.id,
                role: user.role,
                email: user.email,
                borrowed_nums: borrowedNums,
                created_at: user.createdAt,
                updated_at: user.updatedAt,
            });
        }
        const totalUsers = await User.count();
        const pageCount = Math.ceil(totalUsers / pageSize);

        res.status(200).json({
            code: 200,
            users: responseUsers,
            page_count: pageCount,
        });
    } catch (error) {
        console.error(error);
        res.status(500).json({
            code: 500,
            msg: '查询用户失败',
            error: error.message,
        });
    }
};

const getAllHistoriesAdmin = async (req, res) => {
    try {
        const { page = 1, size = 10, search_type, search_target } = req.query;

        const pageNum = parseInt(page, 10);
        const pageSize = parseInt(size, 10);

        const query = {
            where: {},
            include: [
                { model: User, as: 'user', attributes: ['email'] },
                { model: Book, as: 'book', attributes: ['name', 'isbn'] },
            ],
            offset: (pageNum - 1) * pageSize,
            limit: pageSize,
            order: [['createdAt', 'DESC']],
        };

        if (search_target) {
            switch (search_type) {
                case 'email':
                    query.include[0].where = { email: { [Op.like]: `%${search_target}%` } };
                    break;
                case 'name':
                    query.include[1].where = { name: { [Op.like]: `%${search_target}%` } };
                    break;
                case 'isbn':
                    query.include[1].where = { isbn: { [Op.like]: `%${search_target}%` } };
                    break;
            }
        }
        const totalRecords = await History.count(query);
        const pageCount = Math.ceil(totalRecords / pageSize);

        const histories = await History.findAll(query);

        const responseHistories = histories.map((history) => ({
            id: history.id,
            borrow_id: history.borrow_id,
            email: history.user?.email,
            book_name: history.book?.name,
            book_isbn: history.book?.isbn,
            created_at: history.createdAt,
            is_back: history.is_back,
        }));

        res.status(200).json({
            code: 200,
            histories: responseHistories,
            page_count: pageCount,
        });
    } catch (error) {
        console.error(error);
        res.status(500).json({
            code: 500,
            msg: '查询历史失败',
            error: error.message,
        });
    }
};

module.exports = {
    handleAddBookAdmin,
    handleDeleteBookAdmin,
    handleUpdateBookAdmin,
    getAdminSummaryAdmin,
    handleGetAllBooksAdmin,
    handleGetAllUsersAdmin
};