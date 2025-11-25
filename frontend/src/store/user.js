import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import api from '@/services/api';
import router from '@/router';
import { TipoUsuario, isProfissional, isPaciente } from '@/types/usuario.js';
import { parseJwt, getStoredToken, setToken, clearToken } from '@/utils/jwt.js';

export const useUserStore = defineStore('user', () => {
  // estado centralizado do usuario autenticado
  // --- STATE ---
  const user = ref(null);
  const isAuthenticated = ref(!!getStoredToken());

  // --- GETTERS (como computed properties) ---
  const isLoggedIn = computed(() => isAuthenticated.value);
  const userType = computed(() => user.value?.tipo_usuario || null);
  const isProfissionalUser = computed(() => user.value ? isProfissional(user.value) : false);
  const isPacienteUser = computed(() => user.value ? isPaciente(user.value) : false);

  // --- ACTIONS (como funcoes) ---

  /**
   * Busca os dados do perfil do usuario logado paciente ou profissional
   * @param {string} userType - tipo de usuario: 'paciente' ou 'profissional'
   */
  async function fetchUser(userType) {
    // Busca dados apenas se estiver autenticado e perfil ainda nao foi carregado
    if (isAuthenticated.value && !user.value) {
      try {
        let response;
        
        // Usa o enum para comparação type-safe
        if (userType === TipoUsuario.Paciente) {
          response = await api.proprioPerfilPaciente();
        } else if (userType === TipoUsuario.Profissional) {
          response = await api.proprioPerfilProfissional();
        } else {
          throw new Error(`Tipo de usuário inválido: ${userType}`);
        }
        
        user.value = response.data;
      } catch (error) {
        console.error('Erro ao buscar dados do usuário:', error);
        logout();
      }
    }
  }

  async function deleteAccount() {
    try {
      await api.deletarConta();
      logout();
    } catch (error) {
      console.error('Erro ao deletar conta:', error);
      throw error;
    }
  }

  /**
   * Realiza o registro de um novo usuario
   * @param {Object} data - dados do registro
   * @param {string} userType - tipo de usuario: 'paciente' ou 'profissional'
   */
  async function register(data, userType) {
    try {
      let response;
      
      // Usa o enum para comparação type-safe
      if (userType === TipoUsuario.Paciente) {
        response = await api.registrarPaciente(data);
      } else if (userType === TipoUsuario.Profissional) {
        response = await api.registrarProfissional(data);
      } else {
        throw new Error(`Tipo de usuário inválido: ${userType}`);
      }

      // Armazena o token retornado no localStorage
      if (response.data.token) {
        setToken(response.data.token);
        isAuthenticated.value = true;
      }

      return response;
    } catch (error) {
      console.error('Erro ao registrar usuário:', error);
      throw error;
    }
  }

  /**
   * Realiza o login do usuario
   * @param {Object} credentials - credenciais de login email e senha
   */
  async function login(credentials) {
    try {
      const response = await api.login(credentials);
      
      // Armazena o token retornado no localStorage
      if (response.data.token) {
        setToken(response.data.token);
        isAuthenticated.value = true;

        // Decodifica o token para obter o tipo de usuário
        const payload = parseJwt(response.data.token);
        
        // Agora o role vem como string: "profissional" ou "paciente"
        const role = payload?.role;

        // Busca os dados completos do usuário
        if (role) {
          await fetchUser(role);
        }
      }

      return response;
    } catch (error) {
      console.error('Erro ao fazer login:', error);
      throw error;
    }
  }

  /**
   * Realiza o logout do usuario
   */
  function logout() {
    clearToken();
    user.value = null;
    isAuthenticated.value = false;
    router.push('/login');
  }

  // disponibiliza estado e acoes para os componentes
  return {
    user,
    isAuthenticated,
    isLoggedIn,
    userType,
    isProfissionalUser,
    isPacienteUser,
    fetchUser,
    login,
    register,
    logout,
    deleteAccount,
  };
});