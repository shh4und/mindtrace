import { createRouter, createWebHistory } from 'vue-router'
import { TipoUsuario } from '@/types/usuario.js'
import { getUserRoleFromToken, isAuthenticated } from '@/utils/jwt.js'

// Lazy loading dos componentes para code-splitting
const Landpage = () => import('@/views/Landpage.vue')
const Login = () => import('@/views/auth/Login.vue')
const Cadastro = () => import('@/views/auth/Cadastro.vue')
const ForgotPassword = () => import('@/views/auth/ForgotPassword.vue')

// Dashboard Paciente e suas views
const PacienteDashboard = () => import('@/views/dashboard-paciente/PacienteDashboard.vue')
const PacienteResumo = () => import('@/views/dashboard-paciente/Resumo.vue')
const PacienteHumor = () => import('@/views/dashboard-paciente/RegistroHumor.vue')
const PacienteRelatorio = () => import('@/views/shared/Relatorio.vue')
const PacienteVincular = () => import('@/views/dashboard-paciente/VincularProfissional.vue')
const PacienteEditarPerfil = () => import('@/views/shared/EditarPerfil.vue')

// Dashboard Profissional e suas views
const ProfissionalDashboard = () => import('@/views/dashboard-profissional/ProfissionalDashboard.vue')
const ProfissionalPacientes = () => import('@/views/dashboard-profissional/ListaPacientes.vue')
const ProfissionalConvite = () => import('@/views/dashboard-profissional/GerarConvite.vue')
const ProfissionalEditarPerfil = () => import('@/views/shared/EditarPerfil.vue')
const ProfissionalRelatorio = () => import('@/views/shared/Relatorio.vue')

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
    // Dashboard Paciente com nested routes
    {
      path: '/dashboard-paciente',
      name: 'dashboard-paciente',
      component: PacienteDashboard,
      redirect: { name: 'paciente-resumo' },
      meta: { requiresAuth: true, role: TipoUsuario.Paciente },
      children: [
        {
          path: 'resumo',
          name: 'paciente-resumo',
          component: PacienteResumo
        },
        {
          path: 'humor',
          name: 'paciente-humor',
          component: PacienteHumor
        },
        {
          path: 'relatorios',
          name: 'paciente-relatorios',
          component: PacienteRelatorio,
          props: { userType: TipoUsuario.Paciente }
        },
        {
          path: 'vincular',
          name: 'paciente-vincular',
          component: PacienteVincular
        },
        {
          path: 'editar-perfil',
          name: 'paciente-editar-perfil',
          component: PacienteEditarPerfil,
          props: { userType: TipoUsuario.Paciente }
        }
      ]
    },
    // Dashboard Profissional com nested routes
    {
      path: '/dashboard-profissional',
      name: 'dashboard-profissional',
      component: ProfissionalDashboard,
      redirect: { name: 'profissional-pacientes' },
      meta: { requiresAuth: true, role: TipoUsuario.Profissional },
      children: [
        {
          path: 'pacientes',
          name: 'profissional-pacientes',
          component: ProfissionalPacientes
        },
        {
          path: 'pacientes/:patientId/relatorio',
          name: 'profissional-paciente-relatorio',
          component: ProfissionalRelatorio,
          props: route => ({ 
            userType: TipoUsuario.Profissional, 
            patientId: route.params.patientId 
          })
        },
        {
          path: 'convite',
          name: 'profissional-convite',
          component: ProfissionalConvite
        },
        {
          path: 'editar-perfil',
          name: 'profissional-editar-perfil',
          component: ProfissionalEditarPerfil,
          props: { userType: TipoUsuario.Profissional }
        }
      ]
    }
  ]
})

// Guarda de navegação global
router.beforeEach((to, from, next) => {
  const role = getUserRoleFromToken()
  const autenticado = isAuthenticated()

  // Rotas públicas que não precisam de autenticação
  const rotasPublicas = ['landpage', 'login', 'cadastro', 'forgot-password']
  
  // Verifica se é rota pública (incluindo rotas filhas de dashboards)
  const isPublicRoute = rotasPublicas.includes(to.name)
  const isPacienteRoute = to.matched.some(record => record.meta.role === TipoUsuario.Paciente)
  const isProfissionalRoute = to.matched.some(record => record.meta.role === TipoUsuario.Profissional)
  
  if (isPublicRoute) {
    // Se já está logado e tenta acessar página pública, redireciona para dashboard
    if (autenticado && role) {
      if (role === TipoUsuario.Profissional) {
        next({ name: 'profissional-pacientes' })
      } else if (role === TipoUsuario.Paciente) {
        next({ name: 'paciente-resumo' })
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
      if (isProfissionalRoute && role !== TipoUsuario.Profissional) {
        next({ name: 'paciente-resumo' })
      } else if (isPacienteRoute && role !== TipoUsuario.Paciente) {
        next({ name: 'profissional-pacientes' })
      } else {
        next()
      }
    }
  }
})

export default router
