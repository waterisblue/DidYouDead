import { createRouter, createWebHistory } from 'vue-router'
import IndexView  from '@/views/IndexView.vue'
import ChartView  from '@/views/ChartView.vue'
import TableView  from '@/views/TableView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
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
  ]
})

export default router
