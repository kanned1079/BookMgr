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
router.beforeEach((to, from, next) => {
    const userStore = useUserStore();
    const userRole = userStore.thisUser.role;
    if (to.path.startsWith('/admin') && userRole !== 'admin') { // 去往的路径中以/admin开头但是用户不是管理员
        return next('/user/summary');   // 重定向到用户的首页
    }
    next();
});

export default router;