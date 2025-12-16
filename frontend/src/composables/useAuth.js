import { computed } from 'vue';
import { useRouter } from 'vue-router';
import { useUserStore } from '@/store/user';
import { TipoUsuario } from '@/types/usuario.js';
import { getUserRoleFromToken, isAuthenticated, isTokenExpired } from '@/utils/jwt.js';

/**
 * Composable para gerenciamento de autenticação
 * Centraliza lógica de auth usada em múltiplos componentes
 */
export function useAuth() {
  const router = useRouter();
  const userStore = useUserStore();

  // Estado reativo de autenticação
  const isLoggedIn = computed(() => userStore.isLoggedIn);
  const user = computed(() => userStore.user);
  const userType = computed(() => userStore.userType);
  const isProfissional = computed(() => userStore.isProfissionalUser);
  const isPaciente = computed(() => userStore.isPacienteUser);

  /**
   * Obtém o role do usuário atual
   * @returns {string|null}
   */
  function getCurrentRole() {
    return getUserRoleFromToken();
  }

  /**
   * Verifica se o usuário tem permissão para acessar rota
   * @param {string} requiredRole - Role requerido
   * @returns {boolean}
   */
  function hasRole(requiredRole) {
    const currentRole = getCurrentRole();
    return currentRole === requiredRole;
  }

  /**
   * Verifica se a sessão está válida
   * @returns {boolean}
   */
  function isSessionValid() {
    return isAuthenticated() && !isTokenExpired();
  }

  /**
   * Redireciona para dashboard correto baseado no role
   */
  function redirectToDashboard() {
    const role = getCurrentRole();
    
    if (role === TipoUsuario.Profissional) {
      router.push({ name: 'dashboard-profissional' });
    } else if (role === TipoUsuario.Paciente) {
      router.push({ name: 'dashboard-paciente' });
    } else {
      router.push({ name: 'login' });
    }
  }

  /**
   * Realiza login e redireciona
   * @param {Object} credentials - { email, senha }
   * @returns {Promise<boolean>} - true se sucesso
   */
  async function login(credentials) {
    try {
      await userStore.login(credentials);
      redirectToDashboard();
      return true;
    } catch (error) {
      console.error('Erro no login:', error);
      throw error;
    }
  }

  /**
   * Realiza logout e redireciona para login
   */
  function logout() {
    userStore.logout();
  }

  /**
   * Registra novo usuário
   * @param {Object} data - Dados do registro
   * @param {string} userType - 'paciente' ou 'profissional'
   * @returns {Promise<boolean>}
   */
  async function register(data, userType) {
    try {
      await userStore.register(data, userType);
      return true;
    } catch (error) {
      console.error('Erro no registro:', error);
      throw error;
    }
  }

  /**
   * Busca dados do usuário logado
   * @param {string} type - Tipo de usuário
   */
  async function fetchCurrentUser(type = null) {
    const userType = type || getCurrentRole();
    if (userType) {
      await userStore.fetchUser(userType);
    }
  }

  return {
    // Estado
    isLoggedIn,
    user,
    userType,
    isProfissional,
    isPaciente,
    
    // Métodos
    getCurrentRole,
    hasRole,
    isSessionValid,
    redirectToDashboard,
    login,
    logout,
    register,
    fetchCurrentUser
  };
}
