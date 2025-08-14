import { createRouter, createWebHistory } from 'vue-router'
import Landpage from '../views/Landpage.vue'
import Login from '../views/Login.vue'
import Cadastro from '../views/Cadastro.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'landpage',
      component: Landpage
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/cadastro',
      name: 'cadastro',
      component: Cadastro
    }
  ]
})

export default router
