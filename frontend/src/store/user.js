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

  // É importante notar que a ação de LOGIN (que define o token)
  // deve ser tratada no seu componente de Login. Após o login bem-sucedido,
  // você deve salvar o token no localStorage e atualizar `isAuthenticated.value = true;`.

  return {
    user,
    isAuthenticated,
    isLoggedIn,
    fetchUser,
    logout,
  };
});