import { createRouter, createWebHistory } from 'vue-router'
import IndexView  from '@/views/IndexView.vue'
import ChartView  from '@/views/ChartView.vue'
import TableView  from '@/views/TableView.vue'
import LoginView  from '@/views/LoginView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/index',
      name: 'index',
      component: IndexView
    },
    {
      path: '/chart',
      name: 'chart',
      component: ChartView
    },
    {
      path: '/table',
      name: 'table',
      component: TableView
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
  ]
})

export default router
