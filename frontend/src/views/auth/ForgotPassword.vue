<template>
  <div class="min-h-screen bg-gray-50 flex flex-col">
    <!-- Navbar Pública -->
    <NavbarPublic :show-menu="false" />

    <div class="flex flex-grow items-center justify-center px-4">
      <div class="w-full max-w-md">
        <div class="bg-white rounded-2xl shadow-sm border border-gray-200 p-8">
          <h2 class="text-2xl font-semibold text-center text-gray-900 mb-2">Redefinir Senha</h2>

          <div v-if="step === 1" class="space-y-6">
            <p class="text-lg text-center text-gray-500">
              Informe o e-mail cadastrado para receber o código de verificação.
            </p>
            <form @submit.prevent="sendCode">
              <div class="mb-4">
                <label for="email" class="block text-lg font-medium text-gray-700 mb-2">E-mail</label>
                <input type="email" id="email" v-model="email" placeholder="Seu e-mail" class="w-full px-4 py-3 rounded-lg border border-gray-300 transition-colors text-gray-900 placeholder-gray-500" required />
              </div>
              <button type="submit" class="w-full bg-emerald-600 hover:bg-emerald-700 text-white font-medium py-3 px-4 rounded-lg transition-colors duration-200">
                Enviar Código
              </button>
            </form>
          </div>

          <div v-if="step === 2" class="space-y-6">
            <p class="text-lg text-center text-gray-500">
              Insira o código de 4 dígitos enviado para <span class="font-semibold text-gray-900">{{ email }}</span> e sua nova senha.
            </p>
            <form @submit.prevent="resetPassword">
              <div class="mb-4">
                <label class="block text-lg font-medium text-gray-700 mb-2">Código de Verificação</label>
                <div class="flex justify-between space-x-2 md:space-x-4">
                  <input v-for="(item, index) in code" :key="index" type="number" v-model="code[index]" @input="handleCodeInput(index, $event)" @keydown="handleKeyDown(index, $event)" :id="'code-' + (index + 1)" class="w-12 h-12 text-center text-2xl rounded-lg border border-gray-300 bg-gray-50 transition-all focus:border-emerald-400 focus:ring-2 focus:ring-emerald-400/30" maxlength="1" required ref="codeInputRefs" />
                </div>
              </div>

              <div class="mb-4">
                <label for="newPassword" class="block text-lg font-medium text-gray-700 mb-2">Nova Senha</label>
                <input type="password" id="newPassword" v-model="newPassword" placeholder="Nova senha" class="w-full px-4 py-3 rounded-lg border border-gray-300 transition-colors text-gray-900 placeholder-gray-500" required />
              </div>

              <div class="mb-4">
                <label for="confirmPassword" class="block text-lg font-medium text-gray-700 mb-2">Confirmar Nova Senha</label>
                <input type="password" id="confirmPassword" v-model="confirmPassword" placeholder="Confirme a nova senha" class="w-full px-4 py-3 rounded-lg border border-gray-300 transition-colors text-gray-900 placeholder-gray-500" required />
              </div>

              <button type="submit" class="mt-2 w-full bg-emerald-600 hover:bg-emerald-700 text-white font-medium py-3 px-4 rounded-lg transition-colors duration-200">
                Redefinir Senha
              </button>
            </form>
          </div>

          <div class="mt-6 text-center">
            <router-link to="/login" class="block text-lg text-gray-600 hover:text-emerald-600 transition-colors">
              Lembrou da sua senha? Voltar ao login
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { useRouter } from 'vue-router';
import NavbarPublic from '../../components/layout/NavbarPublic.vue';

const router = useRouter();
const step = ref(1);
const email = ref('');
const code = reactive(['', '', '', '']);
const newPassword = ref('');
const confirmPassword = ref('');
const codeInputRefs = ref([]);

const sentCode = '1234'; // Simulação

const sendCode = () => {
  // Simula o envio do código
  alert(`Código de verificação enviado para ${email.value}. Código simulado: ${sentCode}`);
  step.value = 2;
};

const resetPassword = () => {
  const enteredCode = code.join('');
  if (enteredCode !== sentCode) {
    alert('Código de verificação incorreto. Por favor, tente novamente.');
    return;
  }

  if (newPassword.value !== confirmPassword.value) {
    alert('As senhas não coincidem. Por favor, tente novamente.');
    return;
  }

  alert('Senha redefinida com sucesso!');
  router.push('/login');
};

const handleCodeInput = (index, event) => {
  const value = event.target.value;
  if (value.length === 1 && index < code.length - 1) {
    codeInputRefs.value[index + 1].focus();
  }
};

const handleKeyDown = (index, event) => {
  if (event.key === 'Backspace' && code[index] === '' && index > 0) {
    codeInputRefs.value[index - 1].focus();
  }
};
</script>

<style scoped>
/* Remove as setas de inputs de números para navegadores baseados em WebKit (Chrome, Safari, etc.) */
.code-input::-webkit-outer-spin-button,
.code-input::-webkit-inner-spin-button {
  -webkit-appearance: none;
  margin: 0;
}
</style>
