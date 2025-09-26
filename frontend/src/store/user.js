import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import api from '../services/api';
import router from '../router'; // Import router for redirection

export const useUserStore = defineStore('user', () => {
  // --- STATE ---
  const user = ref(null);
  const isAuthenticated = ref(!!localStorage.getItem('token'));

  // --- GETTERS (como computed properties) ---
  const isLoggedIn = computed(() => isAuthenticated.value);

  // --- ACTIONS (como funções) ---

  /**
   * Busca os dados do perfil do usuário logado (paciente ou profissional).
   * @param {string} userType - O tipo de usuário ('paciente' ou 'profissional').
   */
  async function fetchUser(userType) {
    // Busca dados apenas se estiver autenticado e o perfil ainda não foi carregado.
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
        // Se o token for inválido (erro 401/403), desloga o usuário.
        if (error.response && [401, 403].includes(error.response.status)) {
          logout();
        }
      }
    }
  }
  async function deleteAccount() {
      try {
        await api.deletarConta();
        logout(); // Logout após exclusão
        return { success: true };
      } catch (error) {
        console.error('Falha ao deletar conta:', error);
        return { success: false, error: error.response?.data?.erro || 'Erro ao deletar conta' };
      }
    }
  /**
   * Realiza o registro de um novo usuário.
   * @param {Object} data - Dados do registro.
   * @param {string} userType - Tipo de usuário ('paciente' ou 'profissional').
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
   * Realiza o login do usuário.
   * @param {Object} credentials - Credenciais de login (email e senha).
   */
  async function login(credentials) {
    try {
      const response = await api.login(credentials);
      const token = response.data.token;
      localStorage.setItem('token', token);
      isAuthenticated.value = true;
      // Opcional: Buscar dados do usuário após login
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
   * Realiza o logout do usuário.
   */
  function logout() {
    localStorage.removeItem('token');
    user.value = null;
    isAuthenticated.value = false;
    // Redireciona para a página de login para uma experiência de usuário mais fluida.
    router.push('/login');
  }

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