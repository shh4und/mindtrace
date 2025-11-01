<template>
  <div class="min-h-screen bg-gray-50">
  <!-- Navbar publica -->
    <NavbarPublic :show-menu="false" />

  <!-- Container do formulario de login -->
    <div class="flex items-center justify-center px-4 mt-16">
      <div class="w-full max-w-md">
  <!-- Formulario de login -->
        <div class="bg-white rounded-2xl shadow-sm border border-gray-200 p-8">
          <h2 class="text-3xl font-semibold text-center text-gray-900 mb-8">Entrar</h2>

          <form @submit.prevent="handleLogin" class="space-y-6">
            <!-- Campo email -->
            <div>
              <label for="email" class="block text-lg font-medium text-gray-700 mb-2">
                E-mail
              </label>
              <input type="email" id="email" v-model="email" placeholder="Digite seu e-mail"
                class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors text-gray-900 placeholder-gray-500"
                required />
            </div>

            <!-- Campo senha -->
            <div>
              <label for="password" class="block text-lg font-medium text-gray-700 mb-2">
                Senha
              </label>
              <div class="relative">
                <input :type="passwordFieldType" id="password" v-model="password" placeholder="Digite sua senha"
                  class="w-full px-4 py-3 pr-12 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors text-gray-900 placeholder-gray-500"
                  required />
                <button type="button" @click="togglePasswordVisibility"
                  class="absolute right-4 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600 transition-colors">
                  <span v-if="!showPassword" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <font-awesome-icon :icon="faEye"></font-awesome-icon>
                  </span>
                  <span v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <font-awesome-icon :icon="faEyeSlash"></font-awesome-icon>
                  </span>
                </button>
              </div>
            </div>

            <button type="submit"
              class="w-full bg-emerald-600 hover:bg-emerald-700 text-white font-medium py-3 px-4 rounded-lg transition-colors duration-200 focus:ring-2 focus:ring-emerald-500 focus:ring-offset-2 outline-none">
              Entrar
            </button>
          </form>

          <div class="mt-6 text-center space-y-3">
            <router-link to="/recuperar-senha"
              class="block text-lg text-gray-600 hover:text-emerald-600 transition-colors">
              Esqueceu sua senha?
            </router-link>
            <router-link to="/cadastro" class="block text-lg text-gray-600 hover:text-emerald-600 transition-colors">
              Criar conta
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useToast } from 'vue-toastification';
import { useUserStore } from '../../store/user';
import { TipoUsuario } from '../../types/usuario.js';
import { faEye, faEyeSlash } from '@fortawesome/free-solid-svg-icons';
import NavbarPublic from '../../components/layout/NavbarPublic.vue';

const router = useRouter();
const toast = useToast();
const userStore = useUserStore();

const email = ref('');
const password = ref('');
const showPassword = ref(false);

const passwordFieldType = computed(() => (showPassword.value ? 'text' : 'password'));

const togglePasswordVisibility = () => {
  showPassword.value = !showPassword.value;
};

// Funcao para decodificar o payload do jwt
const parseJwt = (token) => {
  try {
    return JSON.parse(atob(token.split('.')[1]));
  } catch (e) {
    return null;
  }
};

const handleLogin = async () => {
  try {
    await userStore.login({ email: email.value, senha: password.value });
    toast.success('Login realizado com sucesso!');
    
    // O redirecionamento é baseado no role do token (agora string)
    const token = localStorage.getItem('token');
    const decodedToken = parseJwt(token);
    
    if (decodedToken && decodedToken.role) {
      // Usa o enum para comparação type-safe
      if (decodedToken.role === TipoUsuario.Profissional) {
        router.push('/dashboard-profissional');
      } else if (decodedToken.role === TipoUsuario.Paciente) {
        router.push('/dashboard-paciente');
      }
    } else {
      toast.error('Token inválido ou tipo de usuário não encontrado.');
    }
  } catch (error) {
    toast.error(error.response?.data?.erro || 'Erro ao fazer login');
  }
};
</script>