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
            <BaseInput
              v-model="email"
              type="email"
              label="E-mail"
              placeholder="Digite seu e-mail"
              autocomplete="email"
              required
              size="lg"
            />

            <!-- Campo senha -->
            <BaseInput
              v-model="password"
              type="password"
              label="Senha"
              placeholder="Digite sua senha"
              autocomplete="current-password"
              required
              size="lg"
            />

            <BaseButton
              type="submit"
              variant="emerald"
              size="lg"
              full-width
            >
              Entrar
            </BaseButton>
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
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useToast } from 'vue-toastification';
import { useUserStore } from '@/store/user';
import { TipoUsuario } from '@/types/usuario.js';
import { parseJwt, getStoredToken } from '@/utils/jwt.js';
import NavbarPublic from '@/components/layout/NavbarPublic.vue';
import { BaseInput, BaseButton } from '@/components/ui';

const router = useRouter();
const toast = useToast();
const userStore = useUserStore();

const email = ref('');
const password = ref('');

const handleLogin = async () => {
  try {
    await userStore.login({ email: email.value, senha: password.value });
    toast.success('Login realizado com sucesso!');
    
    // O redirecionamento é baseado no role do token (agora string)
    const token = getStoredToken();
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