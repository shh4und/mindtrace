<template>
  <div>
    <div class="flex items-center justify-between mb-8">
      <div>
        <h1 class="text-3xl font-bold text-gray-900 mb-2">Questionários Atribuídos</h1>
        <p class="text-gray-500">Acompanhe os questionários atribuídos aos seus pacientes</p>
      </div>
      <button 
        @click="irParaAtribuir" 
        class="px-4 py-2 bg-indigo-600 text-white text-sm font-medium rounded-lg hover:bg-indigo-700 transition-colors flex items-center"
      >
        <font-awesome-icon :icon="faPlus" class="w-4 h-4 mr-2" />
        Atribuir Questionário
      </button>
    </div>

    <!-- Filtros -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-4 mb-6">
      <div class="flex flex-wrap gap-4">
        <div class="flex-1 min-w-[200px]">
          <label class="block text-sm font-medium text-gray-700 mb-1">Filtrar por status</label>
          <select 
            v-model="filtroStatus" 
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
          >
            <option value="">Todos os status</option>
            <option value="PENDENTE">Pendente</option>
            <option value="RESPONDIDO">Respondido</option>
            <option value="EXPIRADO">Expirado</option>
          </select>
        </div>
        <div class="flex-1 min-w-[200px]">
          <label class="block text-sm font-medium text-gray-700 mb-1">Filtrar por paciente</label>
          <input 
            v-model="filtroPaciente" 
            type="text" 
            placeholder="Nome do paciente..."
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
          />
        </div>
        <div class="flex-1 min-w-[200px]">
          <label class="block text-sm font-medium text-gray-700 mb-1">Filtrar por instrumento</label>
          <select 
            v-model="filtroInstrumento" 
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
          >
            <option value="">Todos os instrumentos</option>
            <option value="phq_9">PHQ-9</option>
            <option value="gad_7">GAD-7</option>
            <option value="whoqol_bref">WHOQOL-BREF</option>
            <option value="who_5">WHO-5</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Loading state -->
    <div v-if="isLoading" class="text-center py-12">
      <div class="animate-spin rounded-full h-10 w-10 border-b-2 border-indigo-600 mx-auto mb-4"></div>
      <p class="text-gray-500">Carregando atribuições...</p>
    </div>

    <!-- Empty state -->
    <div v-else-if="atribuicoesFiltradas.length === 0" class="text-center py-12">
      <div class="w-16 h-16 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-4">
        <font-awesome-icon :icon="faClipboardList" class="w-8 h-8 text-gray-400" />
      </div>
      <h3 class="text-lg font-medium text-gray-900 mb-2">
        {{ atribuicoes.length === 0 ? 'Nenhuma atribuição encontrada' : 'Nenhum resultado encontrado' }}
      </h3>
      <p class="text-gray-500">
        {{ atribuicoes.length === 0 
          ? 'Você ainda não atribuiu questionários aos seus pacientes.' 
          : 'Tente ajustar os filtros para encontrar o que procura.' 
        }}
      </p>
    </div>

    <!-- Lista de atribuições -->
    <div v-else class="space-y-4">
      <!-- Estatísticas resumidas -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-4">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-500">Total</p>
              <p class="text-2xl font-bold text-gray-900">{{ atribuicoesFiltradas.length }}</p>
            </div>
            <div class="w-10 h-10 bg-indigo-100 rounded-lg flex items-center justify-center">
              <font-awesome-icon :icon="faClipboardList" class="w-5 h-5 text-indigo-600" />
            </div>
          </div>
        </div>
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-4">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-500">Pendentes</p>
              <p class="text-2xl font-bold text-yellow-600">{{ contarPorStatus('PENDENTE') }}</p>
            </div>
            <div class="w-10 h-10 bg-yellow-100 rounded-lg flex items-center justify-center">
              <font-awesome-icon :icon="faClock" class="w-5 h-5 text-yellow-600" />
            </div>
          </div>
        </div>
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-4">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-500">Respondidos</p>
              <p class="text-2xl font-bold text-emerald-600">{{ contarPorStatus('RESPONDIDO') }}</p>
            </div>
            <div class="w-10 h-10 bg-emerald-100 rounded-lg flex items-center justify-center">
              <font-awesome-icon :icon="faCheckCircle" class="w-5 h-5 text-emerald-600" />
            </div>
          </div>
        </div>
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-4">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-500">Expirados</p>
              <p class="text-2xl font-bold text-gray-600">{{ contarPorStatus('EXPIRADO') }}</p>
            </div>
            <div class="w-10 h-10 bg-gray-100 rounded-lg flex items-center justify-center">
              <font-awesome-icon :icon="faTimesCircle" class="w-5 h-5 text-gray-600" />
            </div>
          </div>
        </div>
      </div>

      <!-- Cards de atribuições -->
      <div 
        v-for="atribuicao in atribuicoesFiltradas" 
        :key="atribuicao.id"
        class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md transition-shadow"
      >
        <div class="flex items-start justify-between">
          <div class="flex-1">
            <!-- Header com paciente -->
            <div class="flex items-center gap-3 mb-3">
              <div class="w-10 h-10 bg-linear-to-br from-indigo-500 to-purple-600 rounded-full flex items-center justify-center text-white font-semibold">
                {{ getIniciais(atribuicao.paciente.nome) }}
              </div>
              <div>
                <h3 class="text-lg font-semibold text-gray-900">{{ atribuicao.paciente.nome }}</h3>
                <p class="text-sm text-gray-500">{{ atribuicao.paciente.email }}</p>
              </div>
            </div>

            <!-- Informações do instrumento -->
            <div class="ml-13 space-y-2">
              <div class="flex items-center gap-3">
                <span 
                  class="inline-block px-2 py-1 text-xs font-mono font-medium rounded-md"
                  :class="getCodigoBadgeClass(atribuicao.instrumento.codigo)"
                >
                  {{ atribuicao.instrumento.codigo.toUpperCase().replace('_', '-') }}
                </span>
                <span class="text-sm font-medium text-gray-700">
                  {{ atribuicao.instrumento.nome }}
                </span>
                <span 
                  class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                  :class="getStatusClass(atribuicao.status)"
                >
                  <span 
                    class="w-1.5 h-1.5 rounded-full mr-1.5" 
                    :class="getStatusDotClass(atribuicao.status)"
                  ></span>
                  {{ getStatusLabel(atribuicao.status) }}
                </span>
              </div>

              <div class="flex items-center text-xs text-gray-500 gap-4">
                <span class="flex items-center">
                  <font-awesome-icon :icon="faCalendar" class="w-3 h-3 mr-1" />
                  Atribuído em {{ formatDate(atribuicao.data_atribuicao) }}
                </span>
                <span 
                  v-if="atribuicao.data_resposta" 
                  class="flex items-center text-emerald-600"
                >
                  <font-awesome-icon :icon="faCheckCircle" class="w-3 h-3 mr-1" />
                  Respondido em {{ formatDate(atribuicao.data_resposta) }}
                </span>
                <span class="flex items-center">
                  <font-awesome-icon :icon="faListOl" class="w-3 h-3 mr-1" />
                  {{ atribuicao.instrumento.total_perguntas }} perguntas
                </span>
              </div>

              <!-- Pontuação se respondido -->
              <div 
                v-if="atribuicao.status === 'RESPONDIDO' && atribuicao.pontuacao !== null"
                class="mt-3 p-3 bg-gray-50 rounded-lg"
              >
                <div class="flex items-center justify-between">
                  <span class="text-sm font-medium text-gray-700">Pontuação obtida:</span>
                  <div class="flex items-center gap-2">
                    <span class="text-lg font-bold text-indigo-600">
                      {{ atribuicao.pontuacao }}
                    </span>
                    <span class="text-sm text-gray-500">
                      / {{ atribuicao.pontuacao_maxima }}
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Ações -->
          <div class="ml-4 flex flex-col gap-2">
            <button 
              v-if="atribuicao.status === 'RESPONDIDO'"
              @click="visualizarRespostas(atribuicao.id)"
              class="px-4 py-2 bg-indigo-600 text-white text-sm font-medium rounded-lg hover:bg-indigo-700 transition-colors flex items-center whitespace-nowrap"
            >
              <font-awesome-icon :icon="faEye" class="w-4 h-4 mr-2" />
              Ver Respostas
            </button>
            <button 
              v-else-if="atribuicao.status === 'PENDENTE'"
              @click="enviarLembrete(atribuicao.id)"
              class="px-4 py-2 bg-amber-600 text-white text-sm font-medium rounded-lg hover:bg-amber-700 transition-colors flex items-center whitespace-nowrap"
            >
              <font-awesome-icon :icon="faBell" class="w-4 h-4 mr-2" />
              Enviar Lembrete
            </button>
            <button 
              @click="confirmarExclusao(atribuicao.id)"
              class="px-4 py-2 bg-red-50 text-red-600 text-sm font-medium rounded-lg hover:bg-red-100 transition-colors flex items-center whitespace-nowrap"
            >
              <font-awesome-icon :icon="faTrash" class="w-4 h-4 mr-2" />
              Excluir
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useToast } from 'vue-toastification';
import { api } from '@/services/api';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import {
  faPlus,
  faClipboardList,
  faClock,
  faCheckCircle,
  faTimesCircle,
  faCalendar,
  faListOl,
  faEye,
  faBell,
  faTrash
} from '@fortawesome/free-solid-svg-icons';

const router = useRouter();
const toast = useToast();

const isLoading = ref(true);
const atribuicoes = ref([]);
const filtroStatus = ref('');
const filtroPaciente = ref('');
const filtroInstrumento = ref('');

const atribuicoesFiltradas = computed(() => {
  let resultado = atribuicoes.value;

  if (filtroStatus.value) {
    resultado = resultado.filter(a => a.status === filtroStatus.value);
  }

  if (filtroPaciente.value) {
    const busca = filtroPaciente.value.toLowerCase();
    resultado = resultado.filter(a => 
      a.paciente.nome.toLowerCase().includes(busca) ||
      a.paciente.email.toLowerCase().includes(busca)
    );
  }

  if (filtroInstrumento.value) {
    resultado = resultado.filter(a => a.instrumento.codigo === filtroInstrumento.value);
  }

  return resultado;
});

const contarPorStatus = (status) => {
  return atribuicoesFiltradas.value.filter(a => a.status === status).length;
};

const getIniciais = (nome) => {
  const partes = nome.split(' ');
  if (partes.length >= 2) {
    return (partes[0][0] + partes[partes.length - 1][0]).toUpperCase();
  }
  return nome.substring(0, 2).toUpperCase();
};

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

const irParaAtribuir = () => {
  router.push({ name: 'profissional-atribuir-questionario' });
};

const visualizarRespostas = (atribuicaoId) => {
  router.push({ 
    name: 'profissional-visualizar-respostas', 
    params: { atribuicaoId } 
  });
};

const enviarLembrete = async (atribuicaoId) => {
  try {
    // TODO: Implementar endpoint de lembrete
    // await api.enviarLembrete(atribuicaoId);
    toast.success('Lembrete enviado ao paciente!');
  } catch (error) {
    toast.error('Erro ao enviar lembrete.');
    console.error(error);
  }
};

const confirmarExclusao = (atribuicaoId) => {
  if (confirm('Tem certeza que deseja excluir esta atribuição? Esta ação não pode ser desfeita.')) {
    excluirAtribuicao(atribuicaoId);
  }
};

const excluirAtribuicao = async (atribuicaoId) => {
  try {
    // TODO: Implementar endpoint de exclusão
    // await api.excluirAtribuicao(atribuicaoId);
    atribuicoes.value = atribuicoes.value.filter(a => a.id !== atribuicaoId);
    toast.success('Atribuição excluída com sucesso!');
  } catch (error) {
    toast.error('Erro ao excluir atribuição.');
    console.error(error);
  }
};

onMounted(async () => {
  try {
    const response = await api.listarAtribuicoesProfissional();
    atribuicoes.value = response.data;
    toast.success('Atribuições carregadas com sucesso!');
  } catch (error) {
    toast.error('Erro ao carregar atribuições.');
    console.error(error);
  } finally {
    isLoading.value = false;
  }
});
</script>
