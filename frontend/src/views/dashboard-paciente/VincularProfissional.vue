<template>
  <div class="flex justify-center items-start pt-10">
    <div class="w-full max-w-lg bg-white rounded-xl shadow-sm border border-gray-200 p-8">
      <h1 class="text-3xl font-bold text-gray-900 mb-2 text-center">Vincular a um Profissional</h1>
      <p class="text-gray-500 mb-8 text-center">Insira o token de convite fornecido pelo seu profissional.</p>

      <form @submit.prevent="submitToken" class="space-y-4">
        <div>
          <label for="token" class="block text-base font-medium text-gray-700 mb-1">Token de Convite</label>
          <input 
            type="text" 
            id="token" 
            v-model="token"
            placeholder="Cole o token aqui"
            class="w-full font-mono tracking-wider px-4 py-3 rounded-lg border border-gray-300 outline-none transition focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20"
            required 
          />
        </div>

        <button 
          type="submit"
          :disabled="isLoading || !token.trim()"
          class="w-full bg-emerald-600 hover:bg-emerald-700 text-white font-medium py-3 px-4 rounded-lg transition-colors duration-200 focus:ring-2 focus:ring-emerald-500 focus:ring-offset-2 outline-none disabled:bg-gray-400 disabled:cursor-not-allowed"
        >
          <span v-if="!isLoading">Vincular Agora</span>
          <span v-else>Vinculando...</span>
        </button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useToast } from 'vue-toastification';
import api from '../../services/api';

const token = ref('');
const isLoading = ref(false);
const toast = useToast();

const submitToken = async () => {
  if (!token.value.trim()) {
    toast.error("Por favor, insira um token.");
    return;
  }

  isLoading.value = true;
  try {
    const response = await api.vincularComToken(token.value.trim());
  toast.success(response.data.mensagem || "Vínculo realizado com sucesso!");
  token.value = ''; // Limpa o campo apos o sucesso
  } catch (error) {
    const errorMessage = error.response?.data?.erro || "Falha ao processar o token.";
    toast.error(errorMessage);
    console.error("Erro ao vincular com token:", error);
  } finally {
    isLoading.value = false;
  }
};
</script>


