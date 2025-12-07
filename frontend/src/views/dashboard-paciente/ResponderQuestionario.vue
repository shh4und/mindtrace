<template>
  <!-- Loading state -->
  <div v-if="carregando" class="max-w-3xl mx-auto pb-8 pt-8">
    <div class="flex flex-col items-center justify-center min-h-[60vh]">
      <div class="animate-spin rounded-full h-16 w-16 border-b-2 border-indigo-600 mb-4"></div>
      <p class="text-gray-600">Carregando questionário...</p>
    </div>
  </div>

  <!-- Error state -->
  <div v-else-if="erro" class="max-w-3xl mx-auto pb-8 pt-8">
    <div class="bg-red-50 border border-red-200 rounded-lg p-6 text-center">
      <font-awesome-icon :icon="faExclamationTriangle" class="w-12 h-12 text-red-500 mb-4" />
      <h2 class="text-xl font-semibold text-red-800 mb-2">Erro ao carregar questionário</h2>
      <p class="text-red-600 mb-4">{{ erro }}</p>
      <button 
        @click="router.push({ name: 'paciente-questionarios' })"
        class="px-6 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 transition-colors"
      >
        Voltar para questionários
      </button>
    </div>
  </div>

  <!-- Content -->
  <div v-else class="max-w-3xl mx-auto pb-8">
    <!-- Header fixo -->
    <div class="sticky top-0 bg-white z-10 pb-4 mb-6 border-b border-gray-200">
      <div class="flex items-center mb-4 pt-4">
        <button 
          @click="confirmarSaida" 
          class="mr-4 p-2 rounded-lg hover:bg-gray-100 transition-colors"
          aria-label="Voltar para questionários"
        >
          <font-awesome-icon :icon="faArrowLeft" class="w-5 h-5 text-gray-600" />
        </button>
        <div class="flex-1">
          <span 
            class="inline-block px-2 py-1 text-xs font-mono font-medium rounded-md mb-1"
            :class="getCodigoBadgeClass(instrumento.codigo)"
          >
            {{ instrumento.codigo?.toUpperCase().replace('_', '-') || 'N/A' }}
          </span>
          <h1 class="text-2xl font-bold text-gray-900">{{ instrumento.nome }}</h1>
        </div>
      </div>

      <!-- Barra de progresso -->
      <div class="bg-gray-100 rounded-full h-2 overflow-hidden">
        <div 
          class="bg-indigo-600 h-2 rounded-full transition-all duration-300"
          :style="{ width: `${progresso}%` }"
        ></div>
      </div>
      <p class="text-sm text-gray-500 mt-2 text-center">
        {{ Object.keys(respostas).length }} de {{ instrumento.perguntas?.length || 0 }} perguntas respondidas
      </p>
    </div>

    <!-- Instrução -->
    <div class="bg-blue-50 border border-blue-200 rounded-lg p-4 mb-6">
      <p class="text-sm text-blue-800">
        <font-awesome-icon :icon="faInfoCircle" class="mr-2" />
        {{ instrumento.instrucao }}
      </p>
    </div>

    <!-- Todas as perguntas em scroll -->
    <div class="space-y-6">
      <div 
        v-for="(pergunta, index) in instrumento.perguntas" 
        :key="pergunta.pergunta_id"
        :id="`pergunta-${pergunta.pergunta_id}`"
        class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 transition-all"
        :class="respostas[pergunta.pergunta_id] !== undefined ? 'border-emerald-200 bg-emerald-50/30' : ''"
      >
        <div class="mb-4">
          <div class="flex items-center gap-2 mb-2">
            <span class="text-xs font-medium text-gray-400 uppercase tracking-wide">
              Pergunta {{ index + 1 }}
            </span>
            <span 
              v-if="respostas[pergunta.pergunta_id] !== undefined"
              class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-emerald-100 text-emerald-700"
            >
              <font-awesome-icon :icon="faCheckCircle" class="w-3 h-3 mr-1" />
              Respondida
            </span>
          </div>
          <p class="text-lg font-medium text-gray-900">
            {{ pergunta.conteudo }}
          </p>
        </div>

        <!-- Opções de resposta (Likert) -->
        <div class="space-y-2">
          <div 
            v-for="opcao in instrumento.opcoes_escala" 
            :key="`${pergunta.pergunta_id}-${opcao.valor}`"
            class="flex items-center p-3 border rounded-lg cursor-pointer transition-all"
            :class="respostas[pergunta.pergunta_id] === opcao.valor 
              ? 'border-indigo-500 bg-indigo-50 ring-2 ring-indigo-500' 
              : 'border-gray-200 hover:border-gray-300 hover:bg-gray-50'"
            @click="selecionarResposta(pergunta.pergunta_id, opcao.valor)"
          >
            <span 
              class="w-8 h-8 rounded-full border-2 flex items-center justify-center mr-3 transition-all shrink-0 text-sm font-semibold"
              :class="respostas[pergunta.pergunta_id] === opcao.valor 
                ? 'border-indigo-500 bg-indigo-500 text-white' 
                : 'border-gray-300 text-gray-500'"
            >
              {{ opcao.valor }}
            </span>
            <span class="flex-1 text-sm text-gray-700">{{ opcao.rotulo }}</span>
            <font-awesome-icon 
              v-if="respostas[pergunta.pergunta_id] === opcao.valor"
              :icon="faCheckCircle" 
              class="w-5 h-5 text-indigo-500 shrink-0" 
            />
          </div>
        </div>
      </div>
    </div>

    <!-- Botão fixo de envio -->
    <div class="sticky bottom-0 bg-white border-t border-gray-200 pt-4 mt-8">
      <button 
        @click="enviarRespostas"
        :disabled="!todasRespondidas"
        class="w-full px-6 py-3 bg-emerald-600 text-white font-medium rounded-lg hover:bg-emerald-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed disabled:bg-gray-400 flex items-center justify-center"
      >
        <font-awesome-icon :icon="faPaperPlane" class="w-5 h-5 mr-2" />
        <span v-if="todasRespondidas">Enviar Respostas</span>
        <span v-else>Responda todas as perguntas para enviar</span>
      </button>
      <p class="text-xs text-gray-500 text-center mt-2">
        {{ Object.keys(respostas).length }}/{{ instrumento.perguntas?.length || 0 }} perguntas respondidas
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { api } from '@/services/api'
import { useRouter, useRoute } from 'vue-router';
import { useToast } from 'vue-toastification';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { 
  faArrowLeft, 
  faInfoCircle, 
  faCheckCircle, 
  faPaperPlane,
  faExclamationTriangle 
} from '@fortawesome/free-solid-svg-icons';

const router = useRouter();
const route = useRoute();
const toast = useToast();

const atribuicaoId = computed(() => route.params.atribuicaoId);
const respostas = ref({});
const carregando = ref(true);
const erro = ref(null);

// Dados do instrumento
const instrumento = ref({
  codigo: '',
  nome: '',
  instrucao: '',
  perguntas: [],
  opcoes_escala: []
});

const progresso = computed(() => {
  const respondidas = Object.keys(respostas.value).length;
  return Math.round((respondidas / instrumento.value.perguntas.length) * 100);
});

const todasRespondidas = computed(() => {
  return instrumento.value.perguntas.every(p => respostas.value[p.pergunta_id] !== undefined);
});

const getCodigoBadgeClass = (codigo) => {
  const classes = {
    'phq_9': 'bg-blue-100 text-blue-700',
    'gad_7': 'bg-amber-100 text-amber-700',
    'whoqol_bref': 'bg-emerald-100 text-emerald-700',
    'who_5': 'bg-purple-100 text-purple-700'
  };
  return classes[codigo] || 'bg-gray-100 text-gray-700';
};

const selecionarResposta = (perguntaId, valor) => {
  console.log('Selecionando:', { perguntaId, valor, respostasAntes: {...respostas.value} });
  respostas.value = {
    ...respostas.value,
    [perguntaId]: valor
  };
  console.log('Respostas depois:', respostas.value);
};

const enviarRespostas = () => {
  // TODO: Integrar com API quando backend estiver pronto
  // const payload = {
  //   atribuicaoId: atribuicaoId.value,
  //   respostas: Object.entries(respostas.value).map(([perguntaId, valor]) => ({
  //     perguntaId: parseInt(perguntaId),
  //     valor
  //   }))
  // };
  // await api.enviarResposta(payload);

  toast.success('Questionário enviado com sucesso! Obrigado por responder.');
  
  setTimeout(() => {
    router.push({ name: 'paciente-questionarios' });
  }, 1500);
};

const confirmarSaida = () => {
  const respondidas = Object.keys(respostas.value).length;
  if (respondidas > 0) {
    if (confirm('Você tem respostas não salvas. Deseja realmente sair?')) {
      router.push({ name: 'paciente-questionarios' });
    }
  } else {
    router.push({ name: 'paciente-questionarios' });
  }
};

onMounted(async () => {
  try {
    carregando.value = true;
    const response = await api.buscarAtribuicao(atribuicaoId.value);
    console.log('Dados recebidos:', response.data);
    
    if (!response.data || !response.data.instrumento) {
      throw new Error('Dados do questionário inválidos');
    }
    
    instrumento.value = response.data.instrumento;
    console.log('Perguntas:', instrumento.value.perguntas);
    console.log('Primeira pergunta:', instrumento.value.perguntas[0]);
    toast.success("Questionário carregado com sucesso!");
  } catch (error) {
    console.error('Erro ao carregar questionário:', error);
    erro.value = error.response?.data?.erro || 'Não foi possível carregar o questionário';
    toast.error(erro.value);
  } finally {
    carregando.value = false;
  }
});
</script>