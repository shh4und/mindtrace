<template>
  <div>
    <!-- Header com botão voltar -->
    <div class="flex items-center mb-6">
      <button 
        @click="voltar" 
        class="mr-4 p-2 rounded-lg hover:bg-gray-100 transition-colors"
        aria-label="Voltar para lista de pacientes"
      >
        <font-awesome-icon :icon="faArrowLeft" class="w-5 h-5 text-gray-600" />
      </button>
      <div>
        <h1 class="text-3xl font-bold text-gray-900">Atribuir Questionário</h1>
        <p class="text-gray-500 mt-1">Selecione um instrumento para atribuir ao paciente</p>
      </div>
    </div>

    <!-- Info do paciente -->
    <div class="bg-indigo-50 border border-indigo-200 rounded-lg p-4 mb-8">
      <div class="flex items-center">
        <div class="w-10 h-10 bg-indigo-500 rounded-full flex items-center justify-center mr-3">
          <font-awesome-icon :icon="faUser" class="w-5 h-5 text-white" />
        </div>
        <div>
          <p class="text-sm text-indigo-600 font-medium">Atribuindo para:</p>
          <p class="text-lg font-semibold text-indigo-900">{{ pacienteNome }}</p>
        </div>
      </div>
    </div>

    <!-- Grid de instrumentos -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <div 
        v-for="instrumento in instrumentos" 
        :key="instrumento.codigo"
        class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-lg hover:border-indigo-300 transition-all duration-200"
      >
        <div class="flex items-start justify-between mb-4">
          <div class="flex-1">
            <span 
              class="inline-block px-2 py-1 text-xs font-mono font-medium rounded-md mb-2"
              :class="getCodigoBadgeClass(instrumento.codigo)"
            >
              {{ instrumento.codigo.toUpperCase().replace('_', '-') }}
            </span>
            <h3 class="text-lg font-semibold text-gray-900">{{ instrumento.nome }}</h3>
          </div>
          <div class="w-12 h-12 rounded-lg flex items-center justify-center" :class="getIconBgClass(instrumento.codigo)">
            <font-awesome-icon :icon="faClipboardList" class="w-6 h-6 text-white" />
          </div>
        </div>
        
        <p class="text-gray-600 text-sm mb-4 line-clamp-3">{{ instrumento.descricao }}</p>
        
        <div class="flex items-center justify-between">
          <span class="text-xs text-gray-400">{{ instrumento.totalPerguntas }} perguntas</span>
          <button 
            @click="atribuirQuestionario(instrumento)"
            class="px-4 py-2 bg-indigo-600 text-white text-sm font-medium rounded-lg hover:bg-indigo-700 transition-colors flex items-center"
          >
            <font-awesome-icon :icon="faPaperPlane" class="w-4 h-4 mr-2" />
            Atribuir
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useToast } from 'vue-toastification';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { 
  faArrowLeft, 
  faUser, 
  faClipboardList, 
  faPaperPlane 
} from '@fortawesome/free-solid-svg-icons';

const router = useRouter();
const route = useRoute();
const toast = useToast();

// ID do paciente vindo da rota
const patientId = computed(() => route.params.patientId);

// Mock: nome do paciente (em produção viria da API ou store)
const pacienteNome = ref('Paciente Selecionado');

// Mock: dados dos instrumentos padronizados
const instrumentos = ref([
  {
    codigo: 'phq_9',
    nome: 'Patient Health Questionnaire-9',
    descricao: 'Instrumento de 9 itens para rastreamento e mensuração da gravidade de sintomas depressivos. Amplamente utilizado em atenção primária e pesquisa clínica.',
    totalPerguntas: 9
  },
  {
    codigo: 'gad_7',
    nome: 'Generalized Anxiety Disorder-7',
    descricao: 'Questionário breve de 7 itens para identificar provável transtorno de ansiedade generalizada e avaliar a gravidade dos sintomas.',
    totalPerguntas: 7
  },
  {
    codigo: 'whoqol_bref',
    nome: 'WHOQOL-BREF',
    descricao: 'Versão abreviada do instrumento de avaliação de qualidade de vida da OMS, contemplando domínios físico, psicológico, relações sociais e meio ambiente.',
    totalPerguntas: 26
  },
  {
    codigo: 'who_5',
    nome: 'WHO-5 Well-Being Index',
    descricao: 'Índice de bem-estar de 5 itens desenvolvido pela OMS para medir o estado subjetivo de bem-estar psicológico nas últimas duas semanas.',
    totalPerguntas: 5
  }
]);

const getCodigoBadgeClass = (codigo) => {
  const classes = {
    'phq_9': 'bg-blue-100 text-blue-700',
    'gad_7': 'bg-amber-100 text-amber-700',
    'whoqol_bref': 'bg-emerald-100 text-emerald-700',
    'who_5': 'bg-purple-100 text-purple-700'
  };
  return classes[codigo] || 'bg-gray-100 text-gray-700';
};

const getIconBgClass = (codigo) => {
  const classes = {
    'phq_9': 'bg-blue-500',
    'gad_7': 'bg-amber-500',
    'whoqol_bref': 'bg-emerald-500',
    'who_5': 'bg-purple-500'
  };
  return classes[codigo] || 'bg-gray-500';
};

const atribuirQuestionario = (instrumento) => {
  // TODO: Integrar com API quando backend estiver pronto
  // await api.atribuirQuestionario(patientId.value, instrumento.codigo);
  
  toast.success(`Questionário "${instrumento.nome}" atribuído com sucesso!`);
  
  // Volta para lista de pacientes após atribuição
  setTimeout(() => {
    router.push({ name: 'profissional-pacientes' });
  }, 1500);
};

const voltar = () => {
  router.push({ name: 'profissional-pacientes' });
};
</script>

<style scoped>
.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
