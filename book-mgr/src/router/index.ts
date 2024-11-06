import { createRouter, createWebHistory } from 'vue-router'
// import HomeView from '../views/HomeView.vue'
import Dashboard from '@/views/DashBoard.vue'
import BookMgr from "@/views/Admin/BookMgr.vue";
import UserMgr from "@/views/Admin/UserMgr.vue";
import AdminSummary from "@/views/Admin/AdminSummary.vue";
//
import Borrow from "@/views/User/Borrow.vue";
import MyBorrowed from "@/views/User/MyBorrowed.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
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
          path: 'borrow',
          component: Borrow,
        },
        {
          path: 'borrowed',
          component: MyBorrowed,
        },
      ],
    }
    // {
    //   path: '/',
    //   name: 'home',
    //   component: Home,
    // },
    // {
    //   path: '/about',
    //   name: 'about',
    //   // route level code-splitting
    //   // this generates a separate chunk (About.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   component: () => import('../views/AboutView.vue'),
    // },

  ],
})

export default router
