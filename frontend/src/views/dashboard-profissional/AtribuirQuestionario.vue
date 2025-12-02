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
          <!-- <span class="text-xs text-gray-400">{{ instrumento.totalPerguntas }} perguntas</span> -->
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
import { ref, computed, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import api from '@/services/api';
import { useToast } from 'vue-toastification';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { 
  faArrowLeft, 
  faUser, 
  faClipboardList, 
  faPaperPlane 
} from '@fortawesome/free-solid-svg-icons';


const props = defineProps({
  patientId: {
    type: [Number, String],
    default: null, // ID do paciente, usado pelo profissional
  },patientNome: {
    type: String,
    default: null,}})

const router = useRouter();
const route = useRoute();
const toast = useToast();

// ID do paciente vindo da rota
// const patientId = computed(() => route.params.patientId);

// Mock: nome do paciente (em produção viria da API ou store)
const pacienteNome = ref('Paciente Selecionado');
const instrumentos = ref([]);

onMounted(async () => {
  try {
    const response = await api.listarQuestionarios()
  
    if (Array.isArray(response.data) && response.data.length > 0){
      const qtdInstrumentos = response.data.length
      instrumentos.value = response.data
      
      toast.success(`${qtdInstrumentos} questionários carregados com sucesso.`)
    }
    else{
      toast.warning(`Houve um problema ao carregar ou não há questionários cadastrados.`)
    }
  } catch (error) {
    toast.error('Erro ao carregar os questionários.');
    console.error(error)
  }
  pacienteNome.value = props.patientNome

})


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
};

const voltar = () => {
  router.push({ name: 'profissional-pacientes' });
};
</script>

<style scoped>
.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
