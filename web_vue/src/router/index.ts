import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ManageCenter from '../views/manageCenter/index.vue'
import ManageBook from '../views/manageCenter/index1.vue'
import Article from '@/views/article.vue'
import Manage from '../views/manage/index.vue'
import Manage1 from '../views/manage/index1.vue'
import Manage2 from '../views/manage/index2.vue'
import ArticleDetail from '@/views/articleDetail.vue'
import Info from '@/views/info.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/manageCenter/selfInfo',
    name: 'manageCenter/selfInfo',
    component: ManageCenter
  },
  {
    path: '/manageCenter/articleManage',
    name: 'manageCenter/articleManage',
    component: ManageBook
  },
  {
    path: '/article',
    name: 'article',
    component: Article
  },
  {
    path:'/manage/articleManage',
    name:'manage/articleManage',
    component: Manage
  },
  {
    path:'/manage/tagManage',
    name:'manage/tagManage',
    component: Manage1
  },
  {
    path:'/manage/userManage',
    name:'manage/userManage',
    component: Manage2
  },
  {
    path:'/articleDetail',
    name:'articleDetail',
    component: ArticleDetail
  },
  {
    path:'/info',
    name:'info',
    component: Info
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
