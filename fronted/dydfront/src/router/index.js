import { createRouter, createWebHistory } from 'vue-router'

import LoginView from '@/views/LoginView.vue'
import BubbleView from '@/views/BubbleView.vue'
import FireServiceView from '@/views/FireServiceView.vue'
import TestamentSaveView from '@/views/TestamentSaveView.vue'
import TechServiceView from '@/views/TechServiceView.vue'


import LoginComponent from '@/components/LoginComponent.vue'
import RegisterComponent from '@/components/RegisterComponent.vue'
import TestamentDesignComponent from '@/components/TestamentDesignComponent.vue'
import MyTestamentComponent from '@/components/MyTestamentComponent.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
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
      component: TestamentSaveView,
      children: [
        {
          path: '/testament/my',
          component: MyTestamentComponent
        },
        {
          path: '/testament/design',
          component: TestamentDesignComponent
        }
      ]
    },
    {
      path: '/tech',
      name: 'tech',
      component: TechServiceView
    }
  ]
})

export default router
