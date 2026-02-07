<template>
  <div class="max-w-6xl mx-auto p-4 md:p-8">
    <header class="mb-10 text-center">
      <div
        class="inline-flex items-center justify-center space-x-2 bg-indigo-50 text-indigo-800 px-4 py-1.5 rounded-full text-sm font-medium mb-4 shadow-sm"
      >
        <font-awesome-icon :icon="faEnvelope" class="h-4 w-4" />
        <span>Convite</span>
      </div>
      <h1
        class="text-3xl md:text-4xl font-extrabold text-gray-900 tracking-tight"
      >
        Convidar Paciente
      </h1>
      <p class="text-gray-500 mt-2 text-lg">
        Gere um token único para um paciente se conectar ao seu perfil.
      </p>
    </header>

    <section class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 max-w-lg mx-auto">
      <div class="space-y-6">
        <div>
          <label for="email" class="block text-sm font-bold text-gray-700 mb-2">
            E-mail para envio de convite
          </label>
          <input
            type="email"
            id="email"
            v-model="emailPac"
            placeholder="Digite o endereço de e-mail"
            class="w-full px-4 py-3.5 rounded-2xl border border-gray-200 bg-gray-50 hover:bg-white focus:ring-2 focus:ring-indigo-500 focus:border-transparent outline-none transition-all text-gray-700 placeholder-gray-400 font-medium"
            required
          />
        </div>

        <button
          @click="generateInvite"
          :disabled="isLoading"
          class="w-full bg-gradient-to-r from-indigo-600 to-violet-600 hover:from-indigo-700 hover:to-violet-700 text-white font-bold py-4 px-6 rounded-2xl shadow-lg hover:shadow-xl hover:-translate-y-0.5 transition-all duration-300 disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:translate-y-0 flex items-center justify-center"
        >
          <span v-if="isLoading" class="flex items-center">
            <svg
              class="animate-spin -ml-1 mr-3 h-5 w-5 text-white"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
            >
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            Gerando...
          </span>
          <span v-else>Gerar Convite</span>
        </button>

        <div v-if="token" class="space-y-4 pt-2">
          <div class="bg-indigo-50 rounded-2xl p-5 border border-indigo-100">
            <div class="flex items-center gap-3 mb-3">
              <div class="p-2 bg-indigo-100 rounded-lg text-indigo-600">
                <font-awesome-icon :icon="faTicket" class="w-5 h-5" />
              </div>
              <h3 class="text-lg font-bold text-indigo-900">Convite Gerado!</h3>
            </div>

            <div class="bg-white rounded-xl p-4 border border-indigo-100">
              <p class="text-xs text-gray-500 font-medium mb-1">Token de Convite:</p>
              <div class="flex items-center justify-between gap-3">
                <span class="font-mono text-lg text-gray-800 break-all font-bold">{{ token }}</span>
                <button
                  @click="copyToken"
                  class="shrink-0 p-2 rounded-lg text-indigo-600 hover:bg-indigo-50 transition-colors"
                  title="Copiar para a área de transferência"
                >
                  <font-awesome-icon :icon="faCopy" class="w-5 h-5" />
                </button>
              </div>
            </div>

            <p class="text-sm text-indigo-700 mt-3 font-medium">
              Válido até:
              <span class="font-bold">{{ expiryDate }}</span>
            </p>
          </div>

          <div
            v-if="copied"
            class="bg-emerald-50 border border-emerald-200 rounded-2xl p-4 text-emerald-700 text-sm font-medium flex items-center gap-2"
          >
            <font-awesome-icon :icon="faCheckCircle" class="w-4 h-4" />
            Token copiado para a área de transferência!
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useClipboard } from '@vueuse/core';
import { useToast } from 'vue-toastification';
import api from '@/services/api';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import {
  faEnvelope,
  faTicket,
  faCopy,
  faCheckCircle,
} from '@fortawesome/free-solid-svg-icons';

const token = ref(null);
const emailPac = ref('');
const expiryDate = ref(null);
const copied = ref(false);
const isLoading = ref(false);

const { copy } = useClipboard({ source: token });
const toast = useToast();

const generateInvite = async () => {
  isLoading.value = true;
  copied.value = false;
  try {
    const response = await api.gerarConvite({ email: emailPac.value });
    token.value = response.data.token;
    const expiry = new Date(response.data.data_expiracao);
    expiryDate.value = expiry.toLocaleString('pt-BR', { dateStyle: 'full', timeStyle: 'short' });
    toast.success('Novo convite gerado!');
  } catch (error) {
    const errorMessage = error.response?.data?.erro || 'Falha ao gerar o convite.';
    toast.error(errorMessage);
    console.error('Erro ao gerar convite:', error);
    token.value = null;
    expiryDate.value = null;
  } finally {
    isLoading.value = false;
  }
};

const copyToken = () => {
  if (!token.value) return;
  copy(token.value);
  toast.success('Token copiado!');
  copied.value = true;
  setTimeout(() => (copied.value = false), 3000);
};
</script>

<style scoped>
/* Mantendo consistência com o estilo global */
</style>
