<template>
  <div class="flex justify-center items-start pt-10">
    <div class="w-full max-w-lg bg-white rounded-xl shadow-sm border border-gray-200 p-8 text-center">
      <h1 class="text-3xl font-bold text-gray-900 mb-2">Convidar Paciente</h1>
      <p class="text-gray-500 mb-8">Gere um token único para um paciente se conectar ao seu perfil.</p>

      <button 
        @click="generateInvite"
        class="w-full bg-indigo-600 hover:bg-indigo-700 text-white font-medium py-3 px-4 rounded-lg transition-colors duration-200 focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 outline-none"
      >
        Gerar Convite
      </button>

      <div v-if="token" class="mt-8">
        <h2 class="text-xl font-semibold text-gray-900 mb-4">Convite Gerado!</h2>
        
        <div class="bg-gray-100 rounded-lg p-4 mb-4">
          <p class="text-sm text-gray-500 mb-1">Token de Convite:</p>
          <div class="relative flex items-center justify-between">
            <span class="font-mono text-xl text-gray-800 break-all pr-12">{{ token }}</span>
            <button @click="copyToken" class="absolute right-2 text-gray-500 hover:text-gray-700 transition-colors" title="Copiar para a área de transferência">
              <i class="fa-regular fa-copy"></i>
            </button>
          </div>
        </div>

        <p class="text-sm text-gray-600">
          Este token é válido até: <span class="font-semibold text-indigo-600">{{ expiryDate }}</span>
        </p>
        
        <div v-if="copied" class="mt-4 p-3 rounded-lg bg-green-100 text-green-800 text-sm">
          Token copiado para a área de transferência!
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useClipboard } from '@vueuse/core'; // Usando uma lib moderna para clipboard
import { useToast } from 'vue-toastification';

const token = ref(null);
const expiryDate = ref(null);
const copied = ref(false);

const { copy } = useClipboard({ source: token });
const toast = useToast();

const generateInvite = () => {
  // Lógica para chamar a API de geração de token viria aqui
  // Por enquanto, vamos simular
  token.value = Math.random().toString(36).substring(2, 12).toUpperCase();
  const expiry = new Date();
  expiry.setHours(expiry.getHours() + 24);
  expiryDate.value = expiry.toLocaleString('pt-BR', { dateStyle: 'medium', timeStyle: 'short' });
  copied.value = false;
};

const copyToken = () => {
  copy(token.value);
  toast.success("Token copiado!");
}
</script>
