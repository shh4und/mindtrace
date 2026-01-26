<template>
  <div class="flex justify-center items-start pt-10">
    <div class="w-full space-y-4 max-w-lg bg-white rounded-xl shadow-sm border border-gray-200 p-8">
      <h1 class="text-3xl font-bold text-gray-900 mb-2 text-center">Vincular a um Profissional</h1>
      <p class="text-gray-500 mb-8 text-center">
        Insira o token de convite fornecido pelo seu profissional clique no link recebido por
        e-mail.
      </p>

      <BaseInput
        v-model="token"
        label="Código do Convite"
        placeholder="Cole o token aqui..."
        :disabled="loading"
        @input="verificarTokenManualmente"
      />
      <div
        v-if="dadosConvite"
        class="bg-emerald-50 p-4 rounded-lg border border-emerald-200 flex items-center gap-3"
      >
        <div class="bg-emerald-100 p-2 rounded-full text-emerald-600">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-6 w-6"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
            />
          </svg>
        </div>
        <div>
          <p class="text-sm text-emerald-800 font-medium">Convite encontrado:</p>
          <p class="text-lg font-bold text-emerald-900">{{ dadosConvite.nomeProfissional }}</p>
          <p class="text-xs text-emerald-700">{{ dadosConvite.especialidade }}</p>
        </div>
      </div>

      <p v-if="erro" class="text-red-500 text-sm mt-1">{{ erro }}</p>
      <BaseButton
        variant="emerald"
        full-width
        :loading="loadingVinculo"
        :disabled="!dadosConvite"
        @click="bindWithToken"
        >Confirmar Vínculo
      </BaseButton>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import api from '@/services/api';
import { useToast } from 'vue-toastification';
import { BaseInput, BaseButton } from '@/components/ui';
import { debounce } from 'lodash';

const route = useRoute();
const router = useRouter();
const toast = useToast();

const token = ref('');
const dadosConvite = ref(null);
const erro = ref('');
const loading = ref(false);
const loadingVinculo = ref(false);

onMounted(() => {
  const tokenUrl = route.query.token;
  if (tokenUrl) {
    token.value = tokenUrl;
    router.replace({ query: null });
  }
});

/**
 * TODO:
 * CONTINUAR DAQUI
 */

// Função que consulta a API para ver de quem é o convite
const buscarInfoToken = async (tokenValue) => {
  if (!tokenValue || tokenValue.length < 10) return; // Tamanho mínimo para evitar chamadas inúteis

  loading.value = true;
  erro.value = '';
  dadosConvite.value = null;

  try {
    // GET /convites/info/:token
    const response = await api.consultarToken(tokenValue);
    dadosConvite.value = response.data;
  } catch (err) {
    erro.value = 'Código inválido ou expirado.';
    dadosConvite.value = null;
  } finally {
    loading.value = false;
  }
};
// Debounce para não chamar a API a cada letra digitada (se for digitação manual)
const verificarTokenManualmente = debounce(() => {
  buscarInfoToken(token.value);
}, 500);
const bindWithToken = async () => {
  if (!token.value.trim()) {
    return;
  }

  loadingVinculo.value = true;
  try {
    await api.vincularComToken(token.value.trim());
    toast.success(`Você agora está vinculado a ${dadosConvite.value.nomeProfissional}`);
    router.push({ name: 'paciente-profissionais' });
  } catch (error) {
    toast.error(err.response?.data?.erro || 'Erro ao vincular.');
    toast.error(errorMessage);
    console.error('Erro ao vincular com token:', error);
  } finally {
    loadingVinculo.value = false;
  }
};
</script>
