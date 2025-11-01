import { createRouter, createWebHistory } from 'vue-router'
import Landpage from '../views/Landpage.vue'
import Login from '../views/auth/Login.vue'
import Cadastro from '../views/auth/Cadastro.vue'
import ForgotPassword from '../views/auth/ForgotPassword.vue'
import PacienteDashboard from '../views/dashboard-paciente/PacienteDashboard.vue'
import ProfissionalDashboard from '../views/dashboard-profissional/ProfissionalDashboard.vue'
import { TipoUsuario } from '../types/usuario.js'

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

/**
 * Extrai o role do token JWT armazenado
 * @returns {string|null} - role do usuário ou null se não houver token
 */
function getUserRoleFromToken() {
  const token = localStorage.getItem('token')
  if (!token) return null
  
  try {
    const payload = JSON.parse(atob(token.split('.')[1]))
    return payload.role // Agora retorna string: "profissional" ou "paciente"
  } catch (error) {
    console.error('Erro ao decodificar token:', error)
    return null
  }
}

/**
 * Verifica se o usuário está autenticado
 * @returns {boolean}
 */
function isAuthenticated() {
  return !!localStorage.getItem('token')
}

// Guarda de navegação global
router.beforeEach((to, from, next) => {
  const role = getUserRoleFromToken()
  const autenticado = isAuthenticated()

  // Rotas públicas que não precisam de autenticação
  const rotasPublicas = ['landpage', 'login', 'cadastro', 'forgot-password']
  
  if (rotasPublicas.includes(to.name)) {
    // Se já está logado e tenta acessar página pública, redireciona para dashboard
    if (autenticado && role) {
      if (role === TipoUsuario.Profissional) {
        next({ name: 'dashboard-profissional' })
      } else if (role === TipoUsuario.Paciente) {
        next({ name: 'dashboard-paciente' })
      } else {
        next()
      }
    } else {
      next()
    }
  } else {
    // Rotas protegidas precisam de autenticação
    if (!autenticado) {
      next({ name: 'login' })
    } else {
      // Verifica se o usuário tem permissão para acessar a rota
      if (to.name === 'dashboard-profissional' && role !== TipoUsuario.Profissional) {
        next({ name: 'dashboard-paciente' })
      } else if (to.name === 'dashboard-paciente' && role !== TipoUsuario.Paciente) {
        next({ name: 'dashboard-profissional' })
      } else {
        next()
      }
    }
  }
})

export default router
