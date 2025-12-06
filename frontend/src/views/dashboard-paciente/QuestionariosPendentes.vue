<template>
  <div>
    <h1 class="text-3xl font-bold text-gray-900 mb-2">Meus Questionários</h1>
    <p class="text-gray-500 mb-8">Responda os questionários atribuídos pelo seu profissional</p>

    <!-- Loading state -->
    <div v-if="isLoading" class="text-center py-12">
      <div class="animate-spin rounded-full h-10 w-10 border-b-2 border-indigo-600 mx-auto mb-4"></div>
      <p class="text-gray-500">Carregando questionários...</p>
    </div>

    <!-- Empty state -->
    <div v-else-if="pendencias.length === 0" class="text-center py-12">
      <div class="w-16 h-16 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-4">
        <font-awesome-icon :icon="faClipboardCheck" class="w-8 h-8 text-gray-400" />
      </div>
      <h3 class="text-lg font-medium text-gray-900 mb-2">Nenhum questionário pendente</h3>
      <p class="text-gray-500">Você não possui questionários para responder no momento.</p>
    </div>

    <!-- Lista de questionários -->
    <div v-else class="space-y-4">
      <div v-for="pendencia in pendencias" :key="pendencia.id"
        class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md transition-shadow">
        <div class="flex items-start justify-between">
          <div class="flex-1">
            <div class="flex items-center gap-3 mb-2">
              <span class="inline-block px-2 py-1 text-xs font-mono font-medium rounded-md"
                :class="getCodigoBadgeClass(pendencia.instrumento.codigo)">
                {{ pendencia.instrumento.codigo.toUpperCase().replace('_', '-') }}
              </span>
              <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                :class="getStatusClass(pendencia.status)">
                <span class="w-1.5 h-1.5 rounded-full mr-1.5" :class="getStatusDotClass(pendencia.status)"></span>
                {{ getStatusLabel(pendencia.status) }}
              </span>
            </div>

            <h3 class="text-lg font-semibold text-gray-900 mb-1">{{ pendencia.instrumento.nome }}</h3>
            <p class="text-sm text-gray-500 mb-3">{{ pendencia.instrumento.descricao }}</p>

            <div class="flex items-center text-xs text-gray-400 gap-4">
              <span class="flex items-center">
                <font-awesome-icon :icon="faCalendar" class="w-3 h-3 mr-1" />
                Atribuído em {{ formatDate(pendencia.data_atribuicao) }}
              </span>
              <span class="flex items-center">
                <font-awesome-icon :icon="faListOl" class="w-3 h-3 mr-1" />
                {{ pendencia.instrumento.total_perguntas }} perguntas
              </span>
            </div>
          </div>

          <div class="ml-4">
            <button v-if="pendencia.status === 'PENDENTE'" @click="responderQuestionario(pendencia.id)"
              class="px-4 py-2 bg-indigo-600 text-white text-sm font-medium rounded-lg hover:bg-indigo-700 transition-colors flex items-center">
              <font-awesome-icon :icon="faPen" class="w-4 h-4 mr-2" />
              Responder
            </button>
            <span v-else-if="pendencia.status === 'RESPONDIDO'"
              class="inline-flex items-center px-4 py-2 text-sm font-medium text-emerald-700 bg-emerald-50 rounded-lg">
              <font-awesome-icon :icon="faCheck" class="w-4 h-4 mr-2" />
              Concluído
            </span>
            <span v-else-if="pendencia.status === 'EXPIRADO'"
              class="inline-flex items-center px-4 py-2 text-sm font-medium text-gray-500 bg-gray-100 rounded-lg">
              <font-awesome-icon :icon="faClock" class="w-4 h-4 mr-2" />
              Expirado
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { api } from '@/services/api'
import { useRouter } from 'vue-router';
import { useToast } from 'vue-toastification';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import {
  faClipboardCheck,
  faCalendar,
  faListOl,
  faPen,
  faCheck,
  faClock
} from '@fortawesome/free-solid-svg-icons';

const router = useRouter();
const isLoading = ref(true);
const toast = useToast();

// Mock: lista de pendências
const pendencias = ref([]);


const getCodigoBadgeClass = (codigo) => {
  const classes = {
    'phq_9': 'bg-blue-100 text-blue-700',
    'gad_7': 'bg-amber-100 text-amber-700',
    'whoqol_bref': 'bg-emerald-100 text-emerald-700',
    'who_5': 'bg-purple-100 text-purple-700'
  };
  return classes[codigo] || 'bg-gray-100 text-gray-700';
};

const getStatusClass = (status) => {
  const classes = {
    'PENDENTE': 'bg-yellow-50 text-yellow-800',
    'RESPONDIDO': 'bg-emerald-50 text-emerald-800',
    'EXPIRADO': 'bg-gray-100 text-gray-600'
  };
  return classes[status] || 'bg-gray-100 text-gray-600';
};

const getStatusDotClass = (status) => {
  const classes = {
    'PENDENTE': 'bg-yellow-500',
    'RESPONDIDO': 'bg-emerald-500',
    'EXPIRADO': 'bg-gray-400'
  };
  return classes[status] || 'bg-gray-400';
};

const getStatusLabel = (status) => {
  const labels = {
    'PENDENTE': 'Pendente',
    'RESPONDIDO': 'Respondido',
    'EXPIRADO': 'Expirado'
  };
  return labels[status] || status;
};

const formatDate = (dateString) => {
  const date = new Date(dateString);
  return date.toLocaleDateString('pt-BR', {
    day: '2-digit',
    month: '2-digit',
    year: 'numeric'
  });
};

const responderQuestionario = (atribuicaoId) => {
  router.push({
    name: 'paciente-responder-questionario',
    params: { atribuicaoId }
  });
};

onMounted(async () => {

  try {
    const response = await api.listarAtribuicoesPaciente();
    pendencias.value = response.data;

    isLoading.value = false;
    toast.success('Seus questionários atribuídos.');

  } catch (error) {
    toast.error('Erro ao carregar seus questionários atribuídos.');
    console.error(error)
    isLoading.value = false;
  }
  
});
</script>
