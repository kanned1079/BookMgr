import {createRouter, createWebHistory} from 'vue-router'
import Dashboard from '@/views/DashBoard.vue'
import BookMgr from "@/views/Admin/BookMgr.vue";
import UserMgr from "@/views/Admin/UserMgr.vue";
import AdminSummary from "@/views/Admin/AdminSummary.vue";
import UserSummary from "@/views/User/UserSummary.vue";
import Borrow from "@/views/User/Borrow.vue";
import MyBorrowed from "@/views/User/MyBorrowed.vue";
import Login from "@/views/Login.vue"
import Register from "@/views/Register.vue"
import BorrowHistory from "@/views/Admin/BorrowHistory.vue";

import {useUserStore} from "@/stores/userinfo";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            redirect: '/login'
        },
        {
            path: '/login',
            component: Login
        },
        {
            path: '/register',
            component: Register
        },
        {
            path: '/admin',
            component: Dashboard,
            children: [
                {
                    path: 'summary',
                    component: AdminSummary
                },
                {
                    path: 'book',
                    component: BookMgr
                },
                {
                    path: 'user',
                    component: UserMgr
                },
                {
                    path: 'log',
                    component: BorrowHistory
                },
            ]
        },
        {
            path: '/user',
            component: Dashboard,
            children: [
                {
                    path: 'summary',
                    component: UserSummary
                },
                {
                    path: 'borrow',
                    component: Borrow
                },
                {
                    path: 'borrowed',
                    component: MyBorrowed
                },
            ],
        }
    ],
});

// 路由拦截器 不允许用户访问管理员界面
// router.beforeEach((to, from, next) => {
//     const userStore = useUserStore();
//     const userRole = userStore.thisUser.role;
//     if (to.path.startsWith('/admin') && userRole !== 'admin') { // 去往的路径中以/admin开头但是用户不是管理员
//         return next('/user/summary');   // 重定向到用户的首页
//     }
//
//     next();
// });

router.beforeEach((to, from, next) => {
    const userStore = useUserStore();

    // 检查用户是否已认证，如果未认证且当前路由不是 /login 或 /register，重定向到 /login
    if (!userStore.authed && to.path !== '/login' && to.path !== '/register') {
        return next({ path: '/login', replace: true }); // 重定向到登录页，不允许返回
    }

    // 如果用户已认证，检查权限
    const userRole = userStore.thisUser.role;
    if (to.path.startsWith('/admin') && userRole !== 'admin') {
        return next({ path: '/user/summary', replace: true }); // 重定向到用户首页，不允许返回
    }

    // 否则允许继续导航
    next();
});

export default router;