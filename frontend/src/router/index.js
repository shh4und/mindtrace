import { createRouter, createWebHistory } from 'vue-router'
import Landpage from '../views/Landpage.vue'
import Login from '../views/auth/Login.vue'
import Cadastro from '../views/auth/Cadastro.vue'
import ForgotPassword from '../views/auth/ForgotPassword.vue'
import PacienteDashboard from '../views/dashboard-paciente/PacienteDashboard.vue'
import ProfissionalDashboard from '../views/dashboard-profissional/ProfissionalDashboard.vue'

// configura roteador principal com historico html5 baseado na base url do ambiente
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  // define rotas declaradas delegando protecao para regras de autenticacao externas
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
    },
    {
      path: '/recuperar-senha',
      name: 'forgot-password',
      component: ForgotPassword
    },
    {
      path: '/dashboard-paciente',
      name: 'dashboard-paciente',
      component: PacienteDashboard
    },
    {
      path: '/dashboard-profissional',
      name: 'dashboard-profissional',
      component: ProfissionalDashboard
    }
  ]
})

export default router
