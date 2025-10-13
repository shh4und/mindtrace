import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import api from '../services/api';
import router from '../router'; // Importa roteador para redirecionamento

export const useUserStore = defineStore('user', () => {
  // estado centralizado do usuario autenticado
  // --- STATE ---
  const user = ref(null);
  const isAuthenticated = ref(!!localStorage.getItem('token'));

  // --- GETTERS (como computed properties) ---
  const isLoggedIn = computed(() => isAuthenticated.value);

  // --- ACTIONS (como funcoes) ---

  /**
   * Busca os dados do perfil do usuario logado paciente ou profissional
   * @param {string} userType - tipo de usuario paciente ou profissional
   */
  async function fetchUser(userType) {
    // Busca dados apenas se estiver autenticado e perfil ainda nao foi carregado
    if (isAuthenticated.value && !user.value) {
      try {
        let response;
        if (userType === 'paciente') {
          response = await api.proprioPerfilPaciente();
        } else if (userType === 'profissional') {
          response = await api.proprioPerfilProfissional();
        } else {
          throw new Error('Tipo de usuário inválido para buscar perfil.');
        }
        user.value = response.data;
      } catch (error) {
        console.error('Falha ao buscar dados do usuário:', error);
        // Se o token for invalido erro 401 403 desloga o usuario
        if (error.response && [401, 403].includes(error.response.status)) {
          logout();
        }
      }
    }
  }
  async function deleteAccount() {
      try {
        await api.deletarConta();
        logout(); // Logout apos exclusao
        return { success: true };
      } catch (error) {
        console.error('Falha ao deletar conta:', error);
        return { success: false, error: error.response?.data?.erro || 'Erro ao deletar conta' };
      }
    }
  /**
   * Realiza o registro de um novo usuario
   * @param {Object} data - dados do registro
   * @param {string} userType - tipo de usuario paciente ou profissional
   */
  async function register(data, userType) {
    try {
      let response;
      if (userType === 'paciente') {
        response = await api.registrarPaciente(data);
      } else if (userType === 'profissional') {
        response = await api.registrarProfissional(data);
      } else {
        throw new Error('Tipo de usuário inválido para registro.');
      }
      return { success: true, message: response.data.mensagem || 'Registro realizado com sucesso!' };
    } catch (error) {
      console.error('Falha no registro:', error);
      console.log('Data:', data)
      return { success: false, error: error.response?.data?.erro || 'Erro no registro' };
    }
  }

  /**
   * Realiza o login do usuario
   * @param {Object} credentials - credenciais de login email e senha
   */
  async function login(credentials) {
    try {
      const response = await api.login(credentials);
      const token = response.data.token;
      localStorage.setItem('token', token);
    isAuthenticated.value = true;
    // Opcional buscar dados do usuario apos login
      const decodedToken = JSON.parse(atob(token.split('.')[1]));
      if (decodedToken && decodedToken.role) {
        await fetchUser(decodedToken.role === 'profissional' ? 'profissional' : 'paciente');
      }
      return { success: true };
    } catch (error) {
      console.error('Falha no login:', error);
      return { success: false, error: error.response?.data?.erro || 'Erro no login' };
    }
  }

  /**
   * Realiza o logout do usuario
   */
  function logout() {
    localStorage.removeItem('token');
    user.value = null;
    isAuthenticated.value = false;
    // Redireciona para pagina de login para uma experiencia de usuario mais fluida
    router.push('/login');
  }

  // disponibiliza estado e acoes para os componentes
  return {
    user,
    isAuthenticated,
    isLoggedIn,
    fetchUser,
    login,
    register,
    logout,
    deleteAccount,
  };
});