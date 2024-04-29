import { createRouter, createWebHistory } from 'vue-router'
import IndexView  from '@/views/IndexView.vue'
import ChartView  from '@/views/ChartView.vue'
import LoginView  from '@/views/LoginView.vue'
import BubbleView from '@/views/BubbleView.vue'
import FireServiceView from '@/views/FireServiceView.vue'
import TestamentSaveView from '@/views/TestamentSaveView.vue'

import LoginComponent from '@/components/LoginComponent.vue'
import RegisterComponent from '@/components/RegisterComponent.vue'

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
      path: '/bubble',
      name: 'bubble',
      component: BubbleView
    },
    {
      path: '/',
      name: 'login',
      component: LoginView,
      children: [
        {
          path: '/',
          name: 'loginComponent',
          component: LoginComponent
        },
        {
          path: '/register',
          name: "registerComponent",
          component: RegisterComponent
        }
      ]
    },
    {
      path: '/fireservice',
      name: 'fireService',
      component: FireServiceView
    },
    {
      path: '/testament',
      name: 'testament',
      component: TestamentSaveView
    }
  ]
})

export default router
